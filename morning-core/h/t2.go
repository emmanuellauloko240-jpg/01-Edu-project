package main

import "strings"

var DigitArt = map[rune][]string{
	'1': {
		"  |  ",
		"  |  ",
		"  |  ",
		"  |  ",
		"  |  ",
	},
	'2': {
		" ___ ",
		"   | ",
		" ___ ",
		" |   ",
		" ___ ",
	},
}

func StringToArt(s string) string {
if s == ""{
	return ""
}
var o string
for _, l := range strings.Split(s,"\n") {
	if l == ""{
		continue
	}
	for i := 0; i < 5; i++ {
		for _, v := range l {
			if _, ok := DigitArt[v]; !ok{
				return ""
			}
			o += DigitArt[v][i]
		}
		o += "\n"
	}
}
return o
}
