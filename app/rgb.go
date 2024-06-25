package app

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type rgb struct {
	currentIndex int
	values       [3]int
}

func newRgb(r int, g int, b int) rgb {
	return rgb{values: [3]int{r, g, b}}
}

func (self *rgb) ScrollCurrentValue(n int) {
	i := self.currentIndex
	self.values[i] = clamp(self.values[i]+n, self.Min()[i], self.Max()[i])
}

func (self *rgb) ScrollValueIndex(n int) {
	self.currentIndex = clamp(self.currentIndex+n, 0, 2)
}

func (self *rgb) CurrentValueIndex() int {
	return self.currentIndex
}

func (self *rgb) WithValue(valueIdx int, value int) input {
	rgb := rgb{values: self.values}
	rgb.values[valueIdx] = value
	return input(&rgb)
}

func (self *rgb) Values() [3]int {
	return self.values
}

func (self *rgb) Max() [3]int {
	return [3]int{255, 255, 255}
}
func (self *rgb) Min() [3]int {
	return [3]int{0, 0, 0}
}
func (self *rgb) Prefix() [3]string {
	return [3]string{"R", "G", "B"}
}

func (self *rgb) ToRgb() rgb {
	return *self
}

var _ input = &rgb{}

func rgbToString(rgb rgb) string {
	r, g, b := rgb.Triple()
	return fmt.Sprintf("rgb(%d %d %d)", r, g, b)
}

func rgbFromString(s string) (rgb, error) {
	var rgb rgb
	var err error

	s = strings.TrimPrefix(s, "rgb(")
	s = strings.ReplaceAll(s, ",", " ")
	s = strings.TrimSuffix(s, ")")
	values := strings.Fields(s)
	if len(values) != 3 {
		return rgb, errors.New("rgb should have 3 arguments")
	}
	rgb.values[0], err = strconv.Atoi(values[0])
	if err != nil {
		return rgb, err
	}
	rgb.values[1], err = strconv.Atoi(values[1])
	if err != nil {
		return rgb, err
	}
	rgb.values[2], err = strconv.Atoi(values[2])
	if err != nil {
		return rgb, err
	}

	min, max := rgb.Min(), rgb.Max()
	for i, v := range rgb.values {
		if v > max[i] || v < min[i] {
			return rgb, fmt.Errorf("value out of range %v", v)
		}
	}

	return rgb, err
}

func (self rgb) inverted() rgb {
	r, g, b := self.Triple()
	return rgb{
		values: [3]int{
			255 - r,
			255 - g,
			255 - b,
		},
	}
}

func (self rgb) Triple() (int, int, int) {
	v := self.values
	return v[0], v[1], v[2]
}
