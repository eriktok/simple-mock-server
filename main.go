package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var dir *string
	var port *string

	dir = flag.String("d", "./json", "stats directory to serve from")
	port = flag.String("p", "8080", "http port to listen on")
	flag.Parse()

	log.Println("Starting stats server with root dir: ", *dir)
	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/", fs)
	log.Println("Listening on: ", *port)
	http.ListenAndServe(":"+*port, nil)
}
