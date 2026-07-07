package main

func GenerateFont() map[rune][]string {
	f := map[rune][]string{}
	for c := rune(32); c <= 126; c++ {
		g := []string{}
		for y := 0; y < 8; y++ {
			s := ""
			for x := 0; x < 8; x++ {
				ch := ' '
				if c != ' ' {
					if (x+y+int(c))%7 == 0 || x == y {
						ch = '*'
					}
					if isVowel(c) && (x-3)*(x-3)+(y-3)*(y-3) < 8 {
						ch = '.'
					}
				}
				s += string(ch)
			}
			g = append(g, s)
		}
		f[c] = g
	}
	return f
}
func isVowel(r rune) bool {
	return r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' || r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}
