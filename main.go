package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var http_dir *string
	var http_port *string

	http_dir = flag.String("d", "/var/tmp/stats", "stats directory to serve from")
	http_port = flag.String("p", "8080", "http port to listen on")
	flag.Parse()

	log.Println("Starting stats server with root dir: ", *http_dir)
	fs := http.FileServer(http.Dir(*http_dir))
	http.Handle("/", fs)
	log.Println("Listening on: ", *http_port)
	http.ListenAndServe(":"+*http_port, nil)
}
