package main

import (
	"log"
	"net/http"
	"time"
)

var (
	ServerAddr = ":1210"
)

func main() {
	// new router
	mux := http.NewServeMux()
	// set routing mode
	mux.HandleFunc("/bye", sayBye)
	server := &http.Server{
		Addr:         ServerAddr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	log.Println("Starting server at" + ServerAddr)
	log.Fatal(server.ListenAndServe())
}

func sayBye(w http.ResponseWriter, r *http.Request) {

	time.Sleep(1 * time.Second)
	_, _ = w.Write([]byte("bye bye from server"))
}
