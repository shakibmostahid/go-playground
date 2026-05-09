package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World from Go! I am a PHP developer learning Go.")
	})

	fmt.Println("Starting Go server on port 8080...")

	http.ListenAndServe(":8080", nil)
}