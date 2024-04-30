package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

const (
	HTTPHeaderAccessControllerAllowOrigin = "Access-Control-Allow-Origin"
	HTTPHeaderContentType                 = "Content-Type"
)

func rootHdler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set(HTTPHeaderAccessControllerAllowOrigin, "*")
	rw.Header().Set(HTTPHeaderContentType, "application/json")

	body := []byte("Hello")
	rw.WriteHeader(http.StatusOK)
	rw.Write(body)
}

func ipaddrHdl(rw http.ResponseWriter, req *http.Request) {

	rw.Header().Set(HTTPHeaderAccessControllerAllowOrigin, "*")
	rw.Header().Set(HTTPHeaderContentType, "application/json")

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		rw.Write([]byte("fail to retrieve address"))
	}

	body := []byte(fmt.Sprintf("Address: %v", addrs))
	rw.WriteHeader(http.StatusOK)
	rw.Write((body))
}

func router(mux *http.ServeMux) {
	mux.HandleFunc("GET /", rootHdler)
	mux.HandleFunc("GET /ipaddr", ipaddrHdl)
}

var port string

func main() {

	flag.StringVar(&port, "port", "9090", "Server port")
	flag.Parse()

	mux := http.NewServeMux()
	router(mux)
	log.Printf("REST server started on port: %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), mux))
}
