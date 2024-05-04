package main

import (
	"flag"
	"fmt"
	"io/fs"
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
		slog.Info(fmt.Sprintf("Unable to read. Reason: %v", err))
		os.Exit(1)
	}
	slog.Info(fmt.Sprintf("Received: %s\n", string(buf[:nr])))
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
		slog.Info(fmt.Sprintf("Unable to start server. Reason: %v", err))
		os.Exit(1)
	}
	defer os.Remove(socket)

	for {
		fd, err := l.Accept()
		if err != nil {
			slog.Info(fmt.Sprintf("Error: %v", err))
			os.Exit(1)
		}
		processConn(fd)
	}
}

var isServer bool

func main() {
	os.Mkdir("/var/socket", fs.ModePerm)
	socketFile := "/var/socket/app.socket"
	flag.BoolVar(&isServer, "server", true, "server mode")
	flag.Parse()

	if isServer {
		err := server(socketFile)
		if err != nil {
			slog.Error(fmt.Sprintf("Server error: %v", err))
		}
	} else {
		count := 0
		mux := http.NewServeMux()
		mux.HandleFunc("GET /", func(rw http.ResponseWriter, req *http.Request) {
			slog.Info(fmt.Sprintf("Request host: %v", req.Host))
			conn, err := net.Dial("unix", socketFile)
			if err != nil {
				slog.Error(fmt.Sprintf("Server error: %v", err))
				rw.WriteHeader(http.StatusExpectationFailed)
				rw.Write([]byte("Fail to connect to socket"))
				return
			}

			count = count + 1
			msg := fmt.Sprintf("%v-Hello", count)
			_, err = conn.Write([]byte(msg))
			if err != nil {
				slog.Error(fmt.Sprintf("Server error: %v", err))
				rw.WriteHeader(http.StatusExpectationFailed)
				rw.Write([]byte("Fail to connect to socket"))
				return
			}
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("Success"))
		})
		port := 9090
		fmt.Printf("REST server started on port: %v", port)
		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), mux)
		if err != nil {
			slog.Error(fmt.Sprintf("Server error: %v", err))
		}
	}
}
