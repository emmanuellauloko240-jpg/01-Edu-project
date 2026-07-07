package main

import (
	"fmt"
	"net/http"
)

func age(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w,"method not allowed", http.StatusMethodNotAllowed)
	}
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Println("Hello, Guest!")

	} else {
		fmt.Println("Hello Alice!")
	}
}
func main()  {
	http.HandleFunc("/hello",age)
	http.ListenAndServe(":8080",nil)
}
