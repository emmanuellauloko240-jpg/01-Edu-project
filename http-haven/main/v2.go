package main

import (
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	switch r.Method {
	case "POST":
		data, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("err")
		}
		defer r.Body.Close()

		if len(data) == 0 {
			fmt.Fprintf(w, "body cannot be empty")
		}
		fmt.Fprintf(w, string(data))
	}

}
