package main

import "fmt"

func main() {
	color1, _ := RGBFromHex("#9c6c23")
	color2, _ := RGBFromHex("#e3d654")

	fmt.Println(color1.Colorize("Hello!"))
	fmt.Println(color2.Colorize("世界!"))
}
