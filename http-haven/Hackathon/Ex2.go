package main

import (
	"fmt"
	"net/http"
)

func PathValidation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprintf(w, "Hello, Guest!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}


