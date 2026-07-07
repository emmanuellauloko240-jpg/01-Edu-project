// package main

// import"fmt"
// func main()  {


// ages := map[string]int{
// 	"Alice": 30,
// 	"Bob":   25,
// 	"Carol": 28,
// }
//  ages["ene"] = 43
//  ages["grace"] = 23

 
// for k, v  := range ages {
// 	if v > 25{
// 		// fmt.Println("found: not corresponding" )
// 		    fmt.Println(k,v)

// 	}else{
// 		fmt.Println(k, "is" ,v ,"which is under 26")
// 	}
// }


// delete(ages,"Bob")
// age, ok := ages["ene"]
// fmt.Println(age, ok)
// }
// 6af424aa4071a5051a1550d724f8c92ae93055fb
package main
package main

import "fmt"

type Employee struct {
	Name     string `json:"name"`
	Age      int	`json:"age"`
	isRemote bool	`json:"name"`
}
type Address struct{
	street string ``
	city string
}
func (a Address) PrintAddress{
	fmt.Println()
}

func (e *Employee) UpdateName(newName string)  {
	e.name = newName
}

func main() {
		address := Address{
		street: "mnmmn",
		city: "Texas",
	}
	employee1 := Employee{
		name : "Alice",
		age : 30,
		isRemote: true,
		Address: address,
	}
	employee1.PrintAddress()


	fmt.Println("Address:", employee1.street, employee1.city)
	employee1.UpdateName("Ella")

	fmt.Println("Employee Name:", employee1.name)
	fmt.Println("Employee Age:", employee1.age)

	job := struct{
		title string
		salary int
	}{
		title: "software engineer",
		salary:  100000,
	}
	fmt.Println("job:", job.title)



}
package main

import "fmt"

type Employee struct {
	Name     string `json:"name"`
	Age      int	`json:"age"`
	isRemote bool	`json:"name"`
}
type Address struct{
	street string ``
	city string
}
func (a Address) PrintAddress{
	fmt.Println()
}

func (e *Employee) UpdateName(newName string)  {
	e.name = newName
}

func main() {
		address := Address{
		street: "mnmmn",
		city: "Texas",
	}
	employee1 := Employee{
		name : "Alice",
		age : 30,
		isRemote: true,
		Address: address,
	}
	employee1.PrintAddress()


	fmt.Println("Address:", employee1.street, employee1.city)
	employee1.UpdateName("Ella")

	fmt.Println("Employee Name:", employee1.name)
	fmt.Println("Employee Age:", employee1.age)

	job := struct{
		title string
		salary int
	}{
		title: "software engineer",
		salary:  100000,
	}
	fmt.Println("job:", job.title)



}
package main

import "fmt"

type Employee struct {
	Name     string `json:"name"`
	Age      int	`json:"age"`
	isRemote bool	`json:"name"`
}
type Address struct{
	street string ``
	city string
}
func (a Address) PrintAddress{
	fmt.Println()
}

func (e *Employee) UpdateName(newName string)  {
	e.name = newName
}

func main() {
		address := Address{
		street: "mnmmn",
		city: "Texas",
	}
	employee1 := Employee{
		name : "Alice",
		age : 30,
		isRemote: true,
		Address: address,
	}
	employee1.PrintAddress()


	fmt.Println("Address:", employee1.street, employee1.city)
	employee1.UpdateName("Ella")

	fmt.Println("Employee Name:", employee1.name)
	fmt.Println("Employee Age:", employee1.age)

	job := struct{
		title string
		salary int
	}{
		title: "software engineer",
		salary:  100000,
	}
	fmt.Println("job:", job.title)



}
