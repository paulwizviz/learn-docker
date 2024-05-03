package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func processConn(c net.Conn) {
	buf := make([]byte, 1024)
	nr, err := c.Read(buf)
	if err != nil {
		log.Fatalf("Unable to read. Reason: %v", err)
	}
	fmt.Printf("Received: %s\n", string(buf[:nr]))
}

func server(socket string) error {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(socket)
		os.Exit(0)
	}()

	// socket listening
	l, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatalf("Unable to start server. Reason: %s", err.Error())
	}
	defer os.Remove(socket)

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		processConn(fd)
	}
}

var isServer bool
var logFile string
var socketFile string

func main() {
	os.Mkdir("/var/socket", fs.ModePerm)
	flag.StringVar(&logFile, "log", "/var/socket/app.log", "name of logfile")
	flag.StringVar(&socketFile, "socket", "/var/socket/app.socket", "socket file")
	flag.BoolVar(&isServer, "server", true, "server mode")
	flag.Parse()

	l, err := os.Create(logFile)
	if err != nil {
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(l, nil))
	slog.SetDefault(logger)

	if isServer {
		err = server(socketFile)
		if err != nil {
			slog.Info("Socket unable of set")
		}
	} else {
		count := 0
		mux := http.NewServeMux()
		mux.HandleFunc("GET /", func(rw http.ResponseWriter, req *http.Request) {
			slog.Info(fmt.Sprintf("Request host: %v", req.Host))
			conn, err := net.Dial("unix", socketFile)
			if err != nil {
				rw.WriteHeader(http.StatusExpectationFailed)
				rw.Write([]byte("Fail to connect to socket"))
				slog.Info("Fail to connect to socket")
				return
			}

			count = count + 1
			msg := fmt.Sprintf("%v-Hello", count)
			_, err = conn.Write([]byte(msg))
			if err != nil {
				rw.WriteHeader(http.StatusExpectationFailed)
				rw.Write([]byte("Fail to connect to socket"))
				slog.Info("Fail to connect to socket")
				return
			}
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("Success"))
		})
		port := 9090
		fmt.Printf("REST server started on port: %v", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), mux))
	}
}
