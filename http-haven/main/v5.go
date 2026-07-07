package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func ExerFive(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "code parameter is required")
	}

	convert, err := strconv.Atoi(code)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "code must be a valid integer")
	}
	if convert < 100 || convert > 599 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "code must be a valid HTTP status code (100–599)")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Responding with status [code]")
	}
}
