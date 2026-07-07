package main

import (
	"fmt"
	"net/http"
)

func page(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("X-Custom-Token")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "X-Custom-Token header is missing")
	}
	fmt.Fprintf(w, "Token recieved: abc123")
	data := r.Header.Get("Content-Type not provided")
	if data == "" {
		fmt.Fprintf(w, "Content-Type not provided")
	}
}
