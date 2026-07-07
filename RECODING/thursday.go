package main

import (
	"fmt"
)

// func CheckNumber(arg string) bool {

// 	for _, v := range arg {
// 		if v >= '0' && v < '9' {
// 			return true
// 		}
// 	}
// 	return false
// }
// func CountAlpha(s string) int {
// 	n := 1
// 	for _, v := range s {
// 		if v >= 'a' && v <= 'z' {
// 			n++
// 		}

// 	}
// 	return n
// }
// func CountChar(str string, c rune) int {
// 	count := 0
// 	for _, v := range str {
// 		if v == c {
// 			count++
// 		}
// 	}
// 	return count

// }
// func PrintIf(str string) string {
// 	if len(str) >= 3 || len(str) == 0 {
// 		return "G\n"
// 	}
// 	return "invalid input \n"

// }
func PrintIfNot(str string) string {
	if len(str) < 3 || len(str) == 0 {
		return "G\n"
	}
	return "Invalid input \n"
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
	// fmt.Println(CheckNumber("Hello"))
	// fmt.Println(CheckNumber("Hello0"))
	// fmt.Println(CountAlpha("Hello world"))
	// fmt.Println(CountAlpha("H e l l o"))
	// fmt.Println(CountAlpha("H1e2l3l4o"))
	// fmt.Println(CountChar("Hello World", 'l'))
	// fmt.Println(CountChar("5  balloons", 5))
	// fmt.Println(CountChar("   ", ' '))
	// fmt.Println(CountChar("The 7 deadly sins", '7'))
	// fmt.Print(PrintIf("abcdefz"))
	// fmt.Print(PrintIf("abc"))
	// fmt.Print(PrintIf(""))
	// fmt.Print(PrintIf("14"))

}


package main

import(
	"fmt"
)
type student struct{
	age int;
	name string
}
func main() {
	john := student {
		age: 30,
		name: "John Jonathan", 
	}

	fmt.Println(john.age)
	var ene student
	ene.name = "ella"
	ene.age = 25
	fmt.Println("name:", ene.name, "and she is", ene.age, "years old")

	// printstudent(ene)


}
// func printstudent(ene student){
// 	fmt.Println("name:", ene.name)
// }

