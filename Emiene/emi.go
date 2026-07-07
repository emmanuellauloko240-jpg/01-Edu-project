package main

import(
	
)

// func CheckNumber(arg string) bool {

// for _, v := range arg {
// 	if v >= '0' && v <= '9'{
// 			return true

// 	}

// }	
// return false
// }	
// func main() {
// 	println(CheckNumber("Hello"))
// 	println(CheckNumber("Hello1"))

// }
func CountAlpha(s string) int {
	count := 0
	for _, v := range s {
		if !(v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z'||){
			count ++
		}
		
	}

	return count
}
func main()  {
	println(CountAlpha("Hell!!?o worl,.d"))
	println(CountAlpha("H e l l o"))
	println(CountAlpha("H1e2l3l4o"))
}

9ce9a0339287e1efba861a3348eb1a6c0bf56d1e