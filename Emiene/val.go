/*
	write a map with a key of rune that print states and value of type string that has the capital of the state
*/
package main

import "fmt"
 
func main() {
	people := map[string]int {
		"Ella": 1234567890,
		"Marvy": 12345667788,
		"Andrew": 12345677890,
	}
	people["CHIOMA"] = 123444444444456

	delete(people, "Andrew")
	for key, _ := range people {
		fmt.Println(key)
	}

}