package main

import (
	"fmt"
	"strings"
)

func makeChar(ch rune) []string {
	if ch >= '0' && ch <= '9' {
		top := ".******."
		empty := "........"
		left := "*......."
		right := ".......*"
		sides := "*......*"
		switch ch {
		case '0':
			return []string{
				top,
				sides,
				sides,
				sides,
				sides,
				sides,
				top,
				empty,
			}
		case '1':
			return []string{
				empty,
				right,
				right,
				right,
				right,
				right,
				right,
				empty,
			}
		case '2':
			return []string{
				top,
				right,
				right,
				top,
				left,
				left,
				top,
				empty,
			}
		case '3':
			return []string{
				top,
				right,
				right,
				top,
				right,
				right,
				top,
				empty,
			}
		case '4':
			return []string{
				sides,
				sides,
				sides,
				top,
				right,
				right,
				right,
				empty,
			}
		case '5':
			return []string{
				top,
				left,
				left,
				top,
				right,
				right,
				top,
				empty,
			}
		case '6':
			return []string{
				top,
				left,
				left,
				top,
				sides,
				sides,
				top,
				empty,
			}
		case '7':
			return []string{
				top,
				right,
				right,
				right,
				right,
				right,
				right,
				empty,
			}
		case '8':
			return []string{
				top,
				sides,
				sides,
				top,
				sides,
				sides,
				top,
				empty,
			}
		case '9':
			return []string{
				top,
				sides,
				sides,
				top,
				right,
				right,
				top,
				empty,
			}
		default:
			return []string{
				empty,
				empty,
				empty,
				empty,
				empty,
				empty,
				empty,
				empty,
			}

		}
	}
	if ch >= 'A' && ch <= 'Z' {
		// top := ".******."
		empty := "........"
		middle:= ".******."
		sides := "*......*"
		isVowel := strings.ContainsRune("AEIOU", ch)
		if isVowel {
			return []string{
				middle,
				sides,
				sides,
				middle,
				sides,
				sides,
				middle,
				empty,

			}
		}else{
			return []string{
			middle,
			sides,
			sides,
			sides,
			sides,
			sides,
			sides,
			middle,
			empty,
		}
		
	}
		// return []string{}
	}
}

func GenerateFont() map[rune][]string {
	line := make(map[rune][]string)
	for i := rune(32); i <= (126); i++ {
		line[i] = makeChar(i)
	}
	return line

}
func main() {
	fmt.Println(GenerateFont())
}
