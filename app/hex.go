package app

import (
	"errors"
	"fmt"
	"strconv"
)

func hexToRgb(s string) (rgb, error) {
	if len(s) != 7 {
		err := errors.New("len of hex != 7")
		return rgb{}, err
	}
	if s[0] != '#' {
		err := errors.New("hex first char is not '#'")
		return rgb{}, err
	}
	s = s[1:]

	number, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return rgb{}, err
	}

	return rgb{
		r: int(number >> 16 & 0xFF),
		g: int(number >> 8 & 0xFF),
		b: int(number & 0xFF),
	}, nil
}

func rgbToHex(rgb rgb) (hex string) {
	return fmt.Sprintf("#%02X%02X%02X", rgb.r, rgb.g, rgb.b)
}
