package app

import "strings"

// color mode for changing sliders in UI
// used as index for inputs field in color struct
type inputMode uint8

const inputModeCount = 2
const (
	rgbInputMode = inputMode(iota)
	hslInputMode
)

type input interface {
	// min values of sliders
	Max() [3]int
	// max values of sliders
	Min() [3]int

	// displayed prefixes for sliders
	Prefix() [3]string

	// get index of the slider that is focused
	CurrentValueIndex() int
	// change which slider is focused
	ScrollValueIndex(n int)

	// values of all sliders
	Values() [3]int
	// change value of focused slider
	ScrollCurrentValue(n int)
	// change value of focused slider to min/max
	ScrollCurrentValueToBound(max bool)
	// construct copy with a slider having specific value
	WithValue(valueIndex int, value int) input

	ToRgb() rgb
}

// parse any color format
// expects bad input (with typos?), sets fields that it can
func parseInput(s string) (inputMode, input, error) {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	switch {
	case strings.HasPrefix(s, "rgb("):
		rgb, err := rgbFromString(s)
		return rgbInputMode, input(&rgb), err
	case strings.HasPrefix(s, "hsl("):
		hsl, err := hslFromString(s)
		return hslInputMode, input(&hsl), err
	default: // hex
		rgb, err := rgbFromHexString(s)
		return rgbInputMode, input(&rgb), err
	}
}
