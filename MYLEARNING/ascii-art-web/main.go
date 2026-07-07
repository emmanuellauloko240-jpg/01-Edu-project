package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
