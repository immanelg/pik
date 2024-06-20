package app

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type hsl struct {
	currentIndex int
	values       [3]int
}

func newHsl(h int, s int, l int) hsl {
	return hsl{values: [3]int{h, s, l}}
}

func (self *hsl) CurrentValueIndex() int {
	return self.currentIndex
}

func (self *hsl) ScrollCurrentValue(n int) {
	i := self.currentIndex
	self.values[i] = clamp(self.values[i]+n, self.Min()[i], self.Max()[i])
}

func (self *hsl) ScrollValueIndex(n int) {
	self.currentIndex = clamp(self.currentIndex+n, 0, 2)
}

func (self *hsl) WithValue(valueIdx int, value int) input {
	hsl := hsl{values: self.values}
	hsl.values[valueIdx] = value
	return input(&hsl)
}

func (self *hsl) Values() [3]int {
	return self.values
}

func (self *hsl) Max() [3]int {
	return [3]int{360, 100, 100}
}

func (self *hsl) Min() [3]int {
	return [3]int{0, 0, 0}
}

func (self *hsl) Prefix() [3]string {
	return [3]string{"H", "S", "L"}
}

func (self *hsl) ToRgb() rgb {
	return hslToRgb(*self)
}

var _ input = &hsl{}

func hslToString(hsl hsl) string {
	h, s, l := hsl.Triple()
	return fmt.Sprintf("hsl(%d %d%% %d%%)", h, s, l)
}

func hslFromString(s string) (hsl hsl, err error) {
	s = strings.TrimPrefix(s, "hsl(")
	s = strings.ReplaceAll(s, ",", " ")
	s = strings.TrimSuffix(s, ")")
	values := strings.Fields(s)
	if len(values) != 3 {
		return hsl, errors.New("hsl should have 3 arguments")
	}
	hsl.values[0], err = strconv.Atoi(strings.TrimSuffix(values[0], "deg"))
	hsl.values[1], err = strconv.Atoi(strings.TrimSuffix(values[1], "%"))
	hsl.values[2], err = strconv.Atoi(strings.TrimSuffix(values[2], "%"))
	return
}

func (self hsl) Triple() (int, int, int) {
	v := self.values
	return v[0], v[1], v[2]
}

func hslToRgb(hsl hsl) rgb {
	v := hsl.Values()
	h := float64(v[0]) / 360.0
	s := float64(v[1]) / 100.0
	l := float64(v[2]) / 100.0

	var r, g, b float64

	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - (l * s)
		}
		p := 2*l - q

		r = hueToRgb(p, q, h+1.0/3.0)
		g = hueToRgb(p, q, h)
		b = hueToRgb(p, q, h-1.0/3.0)
	}

	return rgb{values: [3]int{int(r * 255), int(g * 255), int(b * 255)}}
}

func rgbToHsl(rgb rgb) hsl {
	v := rgb.Values()
	r := float64(v[0]) / 255
	g := float64(v[1]) / 255
	b := float64(v[2]) / 255

	maxVal := max(max(r, g), b)
	minVal := min(min(r, g), b)

	var h, s, l float64

	l = (maxVal + minVal) / 2

	if maxVal == minVal {
		h = 0
		s = 0
	} else {
		d := maxVal - minVal

		if l > 0.5 {
			s = d / (2 - maxVal - minVal)
		} else {
			s = d / (maxVal + minVal)
		}

		switch maxVal {
		case r:
			h = (g - b) / d
			if g < b {
				h += 6
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}

		h *= 60
	}

	return newHsl(int(h), int(s*100), int(l*100))
}

func hueToRgb(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*6*(2.0/3.0-t)
	}
	return p
}
