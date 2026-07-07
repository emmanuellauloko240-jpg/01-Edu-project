package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func BasicMath(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("op")
	a, err := strconv.Atoi(r.URL.Query().Get("a"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "err")

	}

	b, err := strconv.Atoi(r.URL.Query().Get("b"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "err")

	}

	switch name {
	case "add":
		fmt.Fprintf(w, fmt.Sprint(a+b))
	case "subtract":
		fmt.Fprintf(w, fmt.Sprint(a-b))
	case "multiply":
		fmt.Fprintf(w, fmt.Sprint(a*b))
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}
