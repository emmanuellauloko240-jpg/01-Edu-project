package main

import "fmt"

func ValidateBanner(banner map[rune][]string) error{
	if banner == nil{
		return fmt.Errorf("banner is nil")
	}
	if len(banner) != 95{
		return fmt.Errorf("banner has %d entries expected 95",len(banner) )
	}
	for r := rune(32); r <= 126; r++{
		l, ok := banner[r]
		if !ok{
			return fmt.Errorf("missing %c", r)
		}
		if len(l) != 8{
			return fmt.Errorf("character '%c' has %v lines, expected 8",r, len(banner[r]))
		}
	}
	return nil
}
