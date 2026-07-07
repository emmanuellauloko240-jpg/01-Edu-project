package main

import (
	"fmt"
	"net/http"
)

func SecureDashboard(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("X-API-Key")
	if name == "secret123" {
		fmt.Fprintf(w, "Welcome")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

}
