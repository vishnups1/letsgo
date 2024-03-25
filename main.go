package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	srvMux := http.NewServeMux()
	srvMux.HandleFunc("/", index)
	srvMux.HandleFunc("/liveness", liveness)
	srvMux.HandleFunc("/readiness", readiness)

	server := http.Server{
		Handler: srvMux,
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintln("let's go!"))
}

func liveness(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintln("alive!"))
}

func readiness(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintln("ready!"))
}
