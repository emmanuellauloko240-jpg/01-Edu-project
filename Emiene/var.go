package main

import (
)

/* take in string as an input and then return
map with a key type of string and then the value type
of int that will print to dey terminal the name of people and their age*/

// func PadArtRows(rows []string, width int) []string {

// }

// func task(string) map[string]int {
// 	people := map[string] int{
// 	"Ella": 20,
// 	"marvy":  18,
// 	"chioma" : 18,
// 	}
// 	for key, val := range people {
// 		fmt.Println(key,val)
// 	}
// 	return people
// }
// func main()  {
// 	fmt.Println(task("people"))
// } 

// write a function that 
// checks how many times a word
//  appear in a sentence it takes
//  in string and return map[string]int 

// func frequencyWord(s string) map[string]int {

// 	word := strings.Fields(s)
// 	final := make(map[string]int)
// for _, v := range word {
// 	final[v]++
// }
// 	return final
// }
// func main()  {
// 	fmt.Println(frequencyWord("ella is learning go today happy today"))
// }

// write a function that removes  duplicate numbers 

// func removeDuplicate(s []int) []int {
// 	result := []int{}
// 	final := make(map[int]bool)
// 	for _, v := range s {
// 		if final[v] {
// 			continue
// 		}else{
// 			final[v]= true
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }
// write a function that return a numbers that are only duplicate 
// 
func mergeMap(first map[rune][]string, second map[rune][]string) map[rune][]string {
	result := make (map[rune][]string)
	for key, v := range first{
		result[key] = v
	}
	
	for key, v := range second {
		result[key] = v
	}
	return result
} 
func main(){
	// fmt.Println(returnDuplicate([]int{2, 3, 7,7,2, 4, 5, 6, 6}))
}