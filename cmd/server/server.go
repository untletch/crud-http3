package main

import (
	"log"
	"net/http"

	"crud-http3/pkg/server"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	handler := http.DefaultServeMux
	handler.HandleFunc("/", server.GetFunc)

	s := &http3.Server{
		Addr:    ":8000",
		Handler: handler,
	}

	log.Print("About to listen on 8000. Go to https://localhost:8000/")
	log.Fatal(s.ListenAndServeTLS("cert.pem", "key.pem"))
}
