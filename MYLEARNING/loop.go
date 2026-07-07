package main

import(
	"fmt"
)
func main() {
	// for i := 1; i <= 30; i++{
	// 	if i % 3 == 0{
	// 		fmt.Println(i)
	// 	}
	// }
count := 0

for i := 1; i <= 30; i++{
	if i%3 == 0{
	 count++
}
}
fmt.Println(count)
}
