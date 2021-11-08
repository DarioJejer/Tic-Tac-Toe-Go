package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello world")
	})
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}
