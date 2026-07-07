package main

import("fmt")

func main() {
	line := ""
	for r := 0; r <= 5; r++ {
		line += "*"
		line += "."
		line += "*"
	}
	
	fmt.Println(line)
}