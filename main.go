package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Render output for the request
		fmt.Fprintf(w, "Hello, world at %v", time.Now())

		// Log incoming requests
		fmt.Printf("%s %s\n", r.Method, r.URL)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
