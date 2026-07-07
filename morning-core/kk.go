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