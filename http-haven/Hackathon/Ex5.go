package main

import (
	"fmt"
	"net/http"
)

func UserAgent(w http.ResponseWriter, r *http.Request)  {
	name := r.Header.Get("User-Agent")
	if name == ""{
		fmt.Fprintf(w,"err")
	}
	fmt.Fprintf(w,"You are visiting us using: %v",name)
}