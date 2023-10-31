package main

import "fmt"

func main() {
	color1, _ := RGBFromHex("#9c6c23")
	color2, _ := RGBFromHex("#e3d654")

	fmt.Println(color1.Colorize("Hello!"))
	fmt.Println(color2.Colorize("世界!"))

	c1, _ := RGBFromHex("#ff0000")
	c2, _ := RGBFromHex("#00ff00")
	c3, _ := RGBFromHex("#0000ff")

	g := NewGradient(c1, c2, c3)

	fmt.Println(g.Colorize("Hello 世界!"))
}
