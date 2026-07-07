package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	sample := os.Args[1]
	text, err := os.ReadFile(sample)
	if err != nil {
		fmt.Println(err)
	}
	result := string(text)
	line := strings.Split(result, "\n")
	fmt.Println(result)

	word := 'H'
	start := (word - 32) * 9
	height := start + 8
	words := line[start:height]
	for _, v := range words {
		fmt.Println(v)
	}

	// ch := "Hi"
	// text := "Hello"

	// for i, v := range text {
	// 	fmt.Println(i,(string(v)))
	// }
	// fmt.Printf("Total parts:  %d\n", len(word))
	// write a program that takes a string
	//  "Hello" and returns each character and
	// its index

}
