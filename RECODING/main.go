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