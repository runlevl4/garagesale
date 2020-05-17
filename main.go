package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.HandlerFunc(Echo)
	log.Println("Listening on :8000")
	if err := http.ListenAndServe(":8000", h); err != nil {
		log.Fatal(err)
	}
}

// Echo returns a simple response.
func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go service!", r.Method, r.URL.Path)
}
