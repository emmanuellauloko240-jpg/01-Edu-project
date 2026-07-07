package main


func GenerateFont() map[rune][]string {
	f := map[rune][]string{}
	for d := rune(32); d < 126; d++ {
		g := []string{}
		for i := 0; i < 8; i++ {
			h := ""
			for j := 0; j < 8; j++ {
				v := ' '
				if d != ' '{
					if(i+j+int(v)) %7 == 0 || i == j {
						v = '*'
					}
					if isVowel(v) && (i+3)+(i+3) * (j+3)+(j+3) < 8{

					}
					h += string(v)
				}
				g = append(g, h)
			}

		}
		f[d] = g
	}
	return f
}
func isVowel(r rune)bool  {
	return r == 'A'
}