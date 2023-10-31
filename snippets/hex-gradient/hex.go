package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type RGB struct {
	red   uint8
	green uint8
	blue  uint8
}

func (rgb RGB) String() string {
	return fmt.Sprintf("RGB(%v,%v,%v)", rgb.red, rgb.green, rgb.blue)
}

// Equal is used by go-cmp's cmp.Equal
func (rgb RGB) Equal(rgb2 RGB) bool {
	return rgb.red == rgb2.red && rgb.green == rgb2.green && rgb.blue == rgb2.blue
}

func RGBFromHex(s string) (RGB, error) {
	result := RGB{}
	s = strings.ToLower(s)
	if len(s) != 7 || s[0] != '#' {
		return result, errors.New("invalid hex format")
	}

	red, err := strconv.ParseUint(s[1:3], 16, 8)
	if err != nil {
		return result, err
	}
	green, err := strconv.ParseUint(s[3:5], 16, 8)
	if err != nil {
		return result, err
	}
	blue, err := strconv.ParseUint(s[5:7], 16, 8)
	if err != nil {
		return result, err
	}

	return RGB{uint8(red), uint8(green), uint8(blue)}, nil
}

// https://notes.burke.libbey.me/ansi-escape-codes/

const escape = "\x1b"
const controlSequenceIntroducer = escape + "["
const setGraphicsRendition = "m"
const reset = controlSequenceIntroducer + "0" + setGraphicsRendition

func (rgb RGB) Colorize(s string) string {
	args := fmt.Sprintf("38;2;%v;%v;%v", rgb.red, rgb.green, rgb.blue)
	return controlSequenceIntroducer + args + setGraphicsRendition + s + reset
}

func (rgb RGB) Interpolate(rgb2 RGB, proportion float64) RGB {
	if proportion < 0 {
		proportion = 0
	}
	if proportion > 1 {
		proportion = 1
	}

	red := uint8((1-proportion)*float64(rgb.red)) + uint8(proportion*float64(rgb2.red))
	green := uint8((1-proportion)*float64(rgb.green)) + uint8(proportion*float64(rgb2.green))
	blue := uint8((1-proportion)*float64(rgb.blue)) + uint8(proportion*float64(rgb2.blue))

	return RGB{red, green, blue}
}
