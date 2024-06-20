package app

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type rgb struct {
	r int // 0..255
	g int // 0..255
	b int // 0.255
}

type rgba struct {
	rgb
	a int // 0..100
}

func rgbToString(rgb rgb) string {
	return fmt.Sprintf("rgb(%d %d %d)", rgb.r, rgb.g, rgb.b)
}

func rgbFromString(s string) (rgb rgb, err error) {
	s = strings.TrimPrefix(s, "rgb(")
	s = strings.TrimSuffix(s, ")")
	values := strings.Fields(s)
	if len(values) != 3 {
		return rgb, errors.New("rgb should have 3 arguments")
	}
	rgb.r, err = strconv.Atoi(values[0])
	rgb.g, err = strconv.Atoi(values[1])
	rgb.b, err = strconv.Atoi(values[2])
	return
}

func rgbToHexString(rgb rgb) string {
	return fmt.Sprintf("rgb(%d %d %d)", rgb.r, rgb.g, rgb.b)

}

func rgbFromHexString(s string) (rgb rgb, err error) {
	s = strings.TrimPrefix(s, "#")

	number, err := strconv.ParseUint(s, 16, 32)

	rgb.r = int(number >> 16 & 0xFF)
	rgb.g = int(number >> 8 & 0xFF)
	rgb.b = int(number & 0xFF)
	return
}

func (self rgb) inverted() rgb {
	return rgb{
		r: 255 - self.r,
		g: 255 - self.g,
		b: 255 - self.b,
	}
}

func (self rgb) triple() (int, int, int) {
	return self.r, self.g, self.b
}
