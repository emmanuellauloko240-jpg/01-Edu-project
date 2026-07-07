package main

import (
	"fmt"
	"net/http"
)

func ExerFour(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	switch r.Method {
	case "POST":
		r.ParseForm()

		username := r.FormValue("username")
		if username == ""{
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w,"username is required")
		}
		
		language := r.FormValue("language")
		if language == ""{
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w,"language is required")
		}
		fmt.Fprintf(w, "Hello %v, you are coding in %v!",username,language)
	}
}
