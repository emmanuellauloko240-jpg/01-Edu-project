package main

import (
	"fmt"
	"strings"
)

type Animation struct {
	Text   string
	Frames int
	Data   []string
}

func NewAnimation(text string, frames int) *Animation {
	if frames < 1 {
		frames = 1
	}
	return &Animation{Text: text, Frames: frames}
}

func (a *Animation) GenerateSpinFrames() {
	a.Data = GenerateSpin(a.Text, a.Frames)
}

func (a *Animation) GenerateWaveFrames() {
	a.Data = GenerateWave(a.Text, a.Frames)
}

func (a *Animation) GenerateZoomFrames() {
	a.Data = GenerateZoom(a.Text, a.Frames)
}

func (a *Animation) GetFrame(i int) string {
	if len(a.Data) == 0 {
		return ""
	}
	return a.Data[i%len(a.Data)]
}

func (a *Animation) Play() string {
	var builder strings.Builder

	for i, frame := range a.Data {
		builder.WriteString(fmt.Sprintf("=== Frame %d ===\n", i))
		builder.WriteString(frame)
		builder.WriteString("\n\n")
	}

	return builder.String()
}

func GenerateWave(text string, frames int) []string {
	var result []string

	width := 30

	for frame := 0; frame < frames; frame++ {

		lines := make([]string, 10)

		for i := range lines {
			lines[i] = strings.Repeat(" ", width)
		}

		baseRow := 4
		offset := frame % 4

		row := baseRow + offset - 2

		if row < 0 {
			row = 0
		}

		if row > 9 {
			row = 9
		}

		lines[row] = centerText(text, width)

		result = append(result, strings.Join(lines, "\n"))
	}

	return result
}

func GenerateZoom(text string, frames int) []string {
	var result []string

	width := 40

	for frame := 0; frame < frames; frame++ {

		lines := make([]string, 10)

		for i := range lines {
			lines[i] = strings.Repeat(" ", width)
		}

		gap := frame % 5

		zoomed := addSpacing(text, gap)

		lines[4] = centerText(zoomed, width)

		result = append(result, strings.Join(lines, "\n"))
	}

	return result
}

func addSpacing(text string, gap int) string {
	var chars []string
	for _, r := range text {
		chars = append(chars, string(r))
	}
	return strings.Join(chars, strings.Repeat(" ", gap))
}

func GenerateSpin(text string, frames int) []string {
	var result []string

	for f := 0; f < frames; f++ {
		lines := make([]string, 10)

		for i := range lines {
			lines[i] = strings.Repeat(" ", 20)
		}

		switch f % 4 {
		case 0:
			lines[4] = centerText(text, 20)

		case 1:
			for i, ch := range text {
				if i+2 < 10 {
					lines[i+2] = centerText(string(ch), 20)
				}
			}

		case 2:
			lines[4] = centerText(reverse(text), 20)

		case 3:
			r := reverse(text)
			for i, ch := range r {
				if i+2 < 10 {
					lines[i+2] = centerText(string(ch), 20)
				}
			}
		}

		result = append(result, strings.Join(lines, "\n"))
	}

	return result
}

func reverse(text string) string {
	r := []rune(text)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	return strings.Repeat(" ", padding) + text +
		strings.Repeat(" ", width-padding-len(text))
}
6cb90b0158bad6f89d4bd836fe5bc9a2c40a7584