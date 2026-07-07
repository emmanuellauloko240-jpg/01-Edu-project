package main

func GeneratePattern(c rune) []string {
	if c < 'A' || c > 'Z' {
		return []string{}
	}
	if c == 'A' {
		return []string{
			"  ##  ",
			" #  # ",
			" #  # ",
			" #### ",
			" #  # ",
			" #  # ",
			" #  # ",
			"      ",
		}
	}
	if c == 'Z' {
		return []string{
			" #### ",
			"    # ",
			"   #  ",
			"  #   ",
			" #    ",
			" #    ",
			" #### ",
			"      ",
		}
	}
	return []string{
		" #### ",
		" #  # ",
		" #  # ",
		" #  # ",
		" #  # ",
		" #  # ",
		" #### ",
		"      ",
	}
}
