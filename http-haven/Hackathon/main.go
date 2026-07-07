package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", PingPong)
	http.HandleFunc("/hello", PathValidation)
	http.HandleFunc("/count", TextCounter)
	http.HandleFunc("/calculate",BasicMath)
	http.HandleFunc("/agent",UserAgent)
	http.HandleFunc("/dashboard",SecureDashboard)
	http.HandleFunc("/legacy",Legacy)
	http.HandleFunc("/v2",v2)
	fmt.Println("server running at http://localhost:8080")
	http.ListenAndServe(":8080",nil)

}
