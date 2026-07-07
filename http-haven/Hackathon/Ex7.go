package main

import (
	"fmt"
	"net/http"
)

func Legacy(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w,r,"http://localhost:8080/v2",http.StatusMovedPermanently)
}
func v2(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Welcome to version 2" )
}