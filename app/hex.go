package app

import (
	"fmt"
	"strconv"
	"strings"
)

func rgbFromHexString(s string) (rgb rgb, err error) {
	s = strings.TrimPrefix(s, "#")

	number, err := strconv.ParseUint(s, 16, 32)

	rgb.values[0] = int(number >> 16 & 0xFF)
	rgb.values[1] = int(number >> 8 & 0xFF)
	rgb.values[2] = int(number & 0xFF)
	return
}

func rgbToHexString(rgb rgb) (hex string) {
	r, g, b := rgb.Triple()
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}
