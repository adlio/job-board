package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world at %v", time.Now())
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
