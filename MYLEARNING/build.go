package main

import(
	"fmt"
)
func square(n int) int{
	return n*n
}
func main() {
	//  var age int
	// fmt.Println("How old are you?")
	// fmt.Scanln(&age)
	// if age >= 20{
	// 	fmt.Println("you're an adult")
	// }else{
	// 	fmt.Println("you're a baby")
	// }

	for i := 1; i <= 12; i++ {
		result := square(i)
		fmt.Println(i,"Squared is", result)
	}
	
}
