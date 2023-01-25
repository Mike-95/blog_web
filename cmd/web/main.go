package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/post", showPost)
	http.HandleFunc("/post/create", createPost)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}
