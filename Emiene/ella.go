package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter number: ")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")

		} else {
			fmt.Println(i, "odd")

		}
	}
	baf7965a250652b1741b6d089aaf3c6d8bf291b9

}
