package main

import"strings"

func PadArtRows(rows []string, width int) []string{
	result := make([]string,len(rows))
	for i, row := range rows {
		if padding := width-len(row); padding > 0 {
			result[i] = row + strings.Repeat(" ",padding)
		}else {
			result[i] = row
		}
	}
	return result
}
