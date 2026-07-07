package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("usage: go run ")
		return
	}
	text := os.Args[1]
	data, err := os.ReadFile(text)
	if err != nil {
		fmt.Println(err)
	}
	result := string(data)
	line := strings.Split(result, "\n" )

	word := 'H'
	H := (word-32)*9 
	end := H + 8
	lines := line[H:end]
	for _, v := range lines {
		fmt.Println(v)
	}

}
