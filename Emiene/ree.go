// write a map with a key of string that print states
//	and value of type string that has the capital of the state
package main

import "fmt"

func main() {
	ella := map[string]string{
		"Benue": "Makurdi",
		"Edo": "Benin-city",
		"Enugu": "Enugu",
	}
	ella["kogi"] = "kogi"
	delete(ella,"Benue")
	for key, val := range ella{
		fmt.Println(key,val)
	}
}