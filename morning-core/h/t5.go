package main

import"strings"
import"fmt"


type Animation struct{
	text string
	frames int
	data []string
}
func (a *Animation) GenerateSpinFrames()  {
	for i := 0; i < a.frames; i++ {
		a.data = append(a.data, strings.TrimRight(strings.Repeat(fmt.Sprintf("%s%d\n",i),10),"\n\n"))
	}
}
func (a *Animation) GenerateWaveFrames() {
	a.GenerateSpinFrames()
}
func (a *Animation) GenerateZoomFrames()  {
	a.GenerateSpinFrames()
}
func (a *Animation)GetFrame(i int) string   {
	if len(a.data) == 0 {
		return ""
	}
	return a.data[i%len(a.data)]
}
func (a *Animation)Play() string  {
	return strings.Join(a.data,"\n\n")
}