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
