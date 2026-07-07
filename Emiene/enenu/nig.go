package main

import (
	"fmt"
	"os"
	"strings"
)

// word := "Hello"
// start := (word)
// height := start + 8
// words := strings.Split(result, "\n")
// 	words := os.Args[1]
// 	for i, v := range words {
// 		fmt.Println(i,string(v))

// }
func main() {
	sample := os.Args[1]
	text, err := os.ReadFile(sample)
	if err != nil {
		fmt.Println(err)
	}
	result := string(text)
	line := strings.Split(result, "\n")

	word := 'H'
	start := (word - 32) * 9
	end := start + 8
	height := line[start:end]
	for _, v := range height {
		fmt.Println(v)
	}

	text = 'e'
	ed := (text - 32) * 9
	stat := end + 8
	lin := line[ed:stat]

	for _, v := range lin {
		fmt.Println(v)
	}

}
