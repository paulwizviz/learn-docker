package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	URLRootPath  = "/"
	URLAPIPath   = "/api"
	URLHelloPath = "/api/hello"

	HTTPHeaderAccessControllerAllowOrigin = "Access-Control-Allow-Origin"
	HTTPHeaderContentType                 = "Content-Type"
)

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	if req.URL.Path != URLHelloPath {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	body := []byte("Hello")

	rw.Header().Set(HTTPHeaderAccessControllerAllowOrigin, "*")
	rw.Header().Set(HTTPHeaderContentType, "application/json")

	rw.WriteHeader(http.StatusOK)
	rw.Write(body)

}

func runREST(router *mux.Router) {
	router.HandleFunc(URLHelloPath, helloHandler)
	router.Use(mux.CORSMethodMiddleware(router))
}

func main() {
	router := mux.NewRouter()
	runREST(router)
	log.Fatal(http.ListenAndServe("0.0.0.0:9090", router))
}
