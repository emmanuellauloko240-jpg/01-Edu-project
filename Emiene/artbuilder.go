package main

import (
	"fmt"


)

type ArtBuilder struct{
	text string
	style string
}

func NewArtBuilder() *ArtBuilder {
    return &ArtBuilder{
		style: "normal",
	}
}
func (a *ArtBuilder) AddText (text string) * ArtBuilder{
	a.text = text
	return a
}
func (a *ArtBuilder) SetStyle(style string) *ArtBuilder  {
	a.style = style
	return a
}
func (a *ArtBuilder) Build()string  {
	switch a.style{
	case "bold":
		return applyBold(a.text)
	case "italic":
		return applyItalic(a.text)
	case "outline":
		return applyOutline(a.text)
	case "shadow":
		return applyShadow(a.text)
	default:
		return a.text
	}
}

func applyBold(text string) string{
	result := ""
	for _, ch  := range text{
		result += string(ch) + string(ch)
	}
	return  result
}
func applyShadow(text string) string{
	result := ""
	for _, v := range text {
		result +=  string(v) + "_"
	}
	return result
}
func applyItalic(text string)string{
	result := ""
	for _, ch := range text {
		result += "/" + string(ch)
		
	}
	return result
}
func applyOutline(text string) string {
    top := ""
    mid := ""
    bot := ""

    for _, ch := range text {
        	top += "+----+"
         	mid +="| " + string(ch) +" |"
          	bot += "+----+"
    }
	return top + "\n" + mid + "\n" + bot + "\n"

}
func main() {
	result1 := NewArtBuilder().
			SetStyle("bold").
			AddText("Hello").
			Build()
	fmt.Println(result1)
	
		result2 := NewArtBuilder().
			SetStyle("italic").
			AddText("Hello").
			Build()
	fmt.Println(result2)
		result3 := NewArtBuilder().
			SetStyle("outline").
			AddText("Hello").
			Build()
	fmt.Println(result3)
		result4 := NewArtBuilder().
			SetStyle("shadow").
			AddText("Hello").
			Build()
	fmt.Println(result4)
}