package main

import(
	"fmt"
)
func RetainFirstHalf(str string)string {
	if str == ""{
		return ""
	}
	if len(str) == 1 {
		return str
	}
	midpoint := len(str)/2
	return str[0 : midpoint]
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}