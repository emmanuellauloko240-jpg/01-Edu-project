package main

import (
	"strings"
)

func StringToArt(input string)string  {
	hold :=  make(map[rune][]string)
	hold['0']=[] string{
			" ___ ",
			" | | ",
			" | | ",
			" | | ",
			" ___ ", 
		}
	hold['1'] = []string{
		"  |  ",
		"  |  ",
		"  |  ",
		"  |  ",
		"  |  ",
	}
	hold['2'] = []string{
		" ___ ",
		"   ||",
		" ___ ",
		" ||  ",
		" ___ ",
	}
	hold['3'] = []string{
		" ___ ", 
        "    |",
        " ___ ",
        "    |",
        " ___ ",
	}
	hold['4'] = []string{
		" | | ",  
		" | | ",
		" |_| ",
		"   | ",
		"   | ",
	}
	hold['5'] = []string{
		" ___ ",
		" |   ",
		" ___ ",
		"    |",
		" ___ ",
	}
		
	
	line := strings.Split(input,"\n")
	
	for _, v := range line {
		for _, v := range v {
			if v < '0' || v > '9'{
				return ""
			}
		}	
	}
	for i := 0; i < 5; i++ {
	for _, c := range v {
		v += hold[c][i]
	}
}
	return ""
}


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

" __  ",
"   ||",
" __  ",
"||   ",
" ___ ",
},
}

func StringToArt(s string) string {
if s == "" {
return ""
}

var o string
for _, l := range strings.Split(s, "\n") {
if l == "" {
continue
}
for i := 0; i < 5; i++ {
for _, r := range l {
if _, ok := DigitArt[r]; !ok {
return ""
}
o += DigitArt[r][i]
}
o += "\n"
}
}
return o
}
