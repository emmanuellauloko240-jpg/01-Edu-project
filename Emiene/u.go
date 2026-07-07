package main

import "fmt"

// "fmt"


func main() {
nums := []int{1,2,3,4,5,6}
evens := []int{}
for i := 0; i < len(nums); i++ {
	if nums[i]%2 == 0{
		evens = append(evens, nums[i])
	}

}	
	fmt.Println(nums[i])

}