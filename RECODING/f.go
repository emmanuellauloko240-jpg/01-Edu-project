// write a function that checks if a word is the same as the reversed '
// func isPalindrome(s string) bool {}
// convert the word to a rune, loop through for i, j = 0, len(w)-1; i < j; i, j = i+1, j-1 {}
//
//	check if the normal and reverse are thesame if w[i] == w[j] {return true} else return false
package main

import (
	"fmt"
	"strings"
)

// "strings"

// func isPalindrome(s string) bool {
// 	r := strings.ToLower(s)
// 	w := []rune(r)
// 		for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1{
// 			if w[i] == w[j] {
// 				return true
// 			}
// 		}
// 		return false
// }
// func main() {
// 	fmt.Println(isPalindrome("Ma da m"))

// }

// write a function that checks if a number is an even number
//  func isEven(num int) bool {
// 	if num % 2 == 0{
// 		return true
// 	}else{
// 		return false
// 	}
//  }
//  func main(){
// 	fmt.Println(isEven(9))
//  }

// write a function that return the biggest number among pair
// func max(a, b int) int {
// 	if a >= b{
// 		return a
// 	}
// 	return b

// 	}

// func main() {
// fmt.Println(max(15,12))
// }

// write a function that checks a factorial of a number
// hiint -> create a int result with a value of 1 result := 1 use conditional loop to loop through
// for i := 1; i < num; i++ {} and multiply the index
// return you factorial (result)

// func factorial(num int) int {
// 	result := 1
// 	for i := 1; i <= num; i++ {
// 		result *= i
// 	}
// 	return result
// }
// func main()  {
// 	fmt.Println(factorial(5))
// }

// create a function that calculate the area of a function {}
// hint -> multiply the l of area by the w of the area
//  func area(length float64, width float64) float64 {
// 		return length * width
//  }
//  func main()  {
// 	fmt.Println(area(5.5,3.0))
//  }

// write a function that greet user
// usage "hello", name, "!"  -> "hello Ella!"
 func greet(name string) string {
	return  "hello" + " " + name + "!"
 }

 // write a function that count vowel present in a word 
  func countVowel(word string) int {
	count := 0
	w := "aeiou"
	for _, v := range word {
		if strings.ContainsRune(w, v) {
			count++
		}
	}
	return count 
	
  } 
 func main()  {
	fmt.Println(countVowel("ella"))
 }