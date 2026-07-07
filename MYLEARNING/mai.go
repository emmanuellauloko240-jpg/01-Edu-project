package main

import(
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r* http.Request){
	fmt.Fprintf(w,"<h1>Welcome to my ASCII Art Generator!</h1><p>Type some text and turn it into art!</p><a href = '/about'>About</a>")
}
func page(res http.ResponseWriter, req* http.Request) {
	res.Header (). Set("content-type","text/html")
	fmt.Fprintf(res,"<h1>About</h1><p>This app was built by me using Go.</p>")
}

func main()  {
	http.HandleFunc("/",homepage)
	http.HandleFunc("/about",page)
	http.ListenAndServe(":8080",nil)
}