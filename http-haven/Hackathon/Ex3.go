package main

import (
	"fmt"
	"io"
	"net/http"
)

func TextCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Send a POST request with text to count words")
	}else if r.Method == "POST"{

	body, _:= io.ReadAll(r.Body)
		fmt.Fprintf(w,"Character count: %d",len(body))
	}
	
}
