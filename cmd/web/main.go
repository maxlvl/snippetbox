package main

import (
	"log"
	"net/http"
)

func main() {
	// use http.NewServeMux() function to initialize a new servemux
	// then register the home function as the handler for the "/" url pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}