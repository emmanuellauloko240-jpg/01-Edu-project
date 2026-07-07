// package main

// import "strings"

// type ArtBuilder struct{ text, style string }

// func NewArtBuilder() *ArtBuilder { return &ArtBuilder{"", "normal"} }

// func (a *ArtBuilder) AddText(text string) *ArtBuilder { a.text += text; return a }

// func (a *ArtBuilder) SetStyle(s string) *ArtBuilder {
// if s != "normal" && s != "bold" && s != "italic" && s != "outline" {
// panic("invalid style")
// }

// a.style = s
// return a
// }

// func (a *ArtBuilder) Build() string {
// r := a.text

// if a.style == "bold" {
// r += a.text
// }
// if a.style == "italic" {
// r = "/" + r
// }
// if a.style == "outline" {
// r = "[" + r + "]"
// }

// return strings.Repeat(r+"\n", 8)
// }
package main