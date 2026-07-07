// write a function that checks if a word is the same as the reversed '
// func isPalindrome(s string) bool {}
// convert the word to a rune, loop through for i, j = 0, len(w)-1; i < j; i, j = i+1, j-1 {}
//  check if the normal and reverse are thesame if w[i] == w[j] {return true} else return false 
// package main

// import "fmt"

// func greet(name string, age int, class string, state string) string {
// 	fmt.Printf("hello my name is %s i am %d years old iam in %v from %s ", name,age,class,state)
// 	return ""
// }
// func main(){
// 	fmt.Println(greet("Emiene", 21,"ss3", "benue"))
// }


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


func isPalindrome(s string) bool {
	r := strings.ToLower(s)
	w := []rune(r)
		for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1{
			if w[i] == w[j] {
				return true
			}
		}
		return false
}
func main() {
	fmt.Println(isPalindrome("Ma da m"))

}
