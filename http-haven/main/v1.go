package main

import (
	"fmt"
	"net/http"
)

func homepagee(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "You made a %v request", r.Method)
	case "POST":
		fmt.Fprintf(w, "you made a %v request", r.Method)
	default:
		fmt.Fprintf(w, "you made a %v request", r.Method)
	}
}


