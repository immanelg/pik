package main

import (
	"errors"
	"fmt"
	"strconv"
)

func hexToRgb(s string) (rgb rgb, err error) {
	if len(s) != 7 {
		err = errors.New("len of hex != 7")
		return
	}
	if s[0] != '#' {
		err = errors.New("hex first char is not '#'")
		return
	}
	r, err := strconv.ParseInt(string(s[1:3]), 16, 16)
	if err != nil {
		return
	}
	g, err := strconv.ParseInt(string(s[3:5]), 16, 16)
	if err != nil {
		return
	}
	b, err := strconv.ParseInt(string(s[5:7]), 16, 16)
	if err != nil {
		return
	}
	rgb.r, rgb.g, rgb.b = r, g, b
	return
}

func rgbToHex(rgb rgb) (s string) {
	return fmt.Sprintf("#%02X%02X%02X", rgb.r, rgb.g, rgb.b)
}
