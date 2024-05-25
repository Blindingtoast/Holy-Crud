package main

import (
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./Golang/static/main.html")
}

// Make a create endpoint
func createPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./Golang/static/create.html")
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/create", createPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
