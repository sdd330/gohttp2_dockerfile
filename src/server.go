package main

import (
	"log"
	"net/http"

	"github.com/bradfitz/http2"
)

func main() {
	var srv http.Server
	srv.Addr = ":443"
	http2.ConfigureServer(&srv, &http2.Server{})
	http.Handle("/", http.FileServer(http.Dir("/www")))
	log.Fatal(srv.ListenAndServeTLS("/app/certs/server.crt", "/app/certs/server.key"))
}