package main

import (
	"math"
	"strings"
)

type Gradient struct {
	colors []RGB
}

func NewGradient(colors ...RGB) Gradient {
	return Gradient{colors}
}

func (g Gradient) Colorize(s string) string {
	builder := strings.Builder{}

	stringLen := len(s)
	colorsLen := len(g.colors)
	for i, c := range s {
		continuous_i := float64(i) / float64(stringLen-1)

		left_rgb_i := int(math.Floor(continuous_i * float64(colorsLen-1)))
		right_rgb_i := left_rgb_i + 1
		if right_rgb_i == colorsLen {
			right_rgb_i -= 1
		}

		left_rgb := g.colors[left_rgb_i]
		right_rgb := g.colors[right_rgb_i]

		prop := proportion(float64(left_rgb_i)/float64(colorsLen-1), continuous_i, float64(right_rgb_i)/float64(colorsLen-1))

		rgb := left_rgb.Interpolate(right_rgb, prop)

		_, _ = builder.WriteString(rgb.Colorize(string(c)))
	}

	return builder.String()
}

func proportion(left, middle, right float64) float64 {
	if left == right {
		return 1
	}

	return (middle - left) / (right - left)
}
