package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from CodeVault"))
}

func codeView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific code"))
}

func codeCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new code"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/code/view", codeView)
	mux.HandleFunc("/code/create", codeCreate)
	
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}