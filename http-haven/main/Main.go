package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/method-inspector", homepagee)
	http.HandleFunc("/echo", home)
	http.HandleFunc("/headers", page)
	http.HandleFunc("/form", ExerFour)
	http.HandleFunc("/status",ExerFive)

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
