package main

import (
	"log"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

func main() {
	addr := "127.0.0.1:8081"
	router := httptreemux.NewContextMux()

	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
