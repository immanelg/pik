package app

import (
	"log"
	"strings"
)

// NOTE: hsl<->rgb<->whatever is not a bijection.
// We can't have only one piece of state for color and then have different "views" to them (rgb/hsl/hsv/whatever).
// When changing a hsl slider value in this model, we have to convert rgb->hsl on the fly, change hsl, and then
// convert it back; it will lose information.

// Displayed color mode
type inputMode uint8

const (
	rgbInputMode inputMode = 0
	hslInputMode           = 1
)

// What will be printed after exiting
type outputMode uint8

const (
	hexOutputMode = outputMode(iota)
	rgbOutputMode
	hslOutputMode
)

type input interface {
	// constraints for values
	Max() [3]int
	Min() [3]int

	// displayed prefixes for sliders
	Prefix() [3]string

	// which slider is focused
	CurrentValueIndex() int
	ScrollValueIndex(n int)

	// values for focused slider
	Values() [3]int
	ScrollCurrentValue(n int)
	WithValue(valueIndex int, value int) input

	ToRgb() rgb
}

// manages state of input and output
type color struct {
	outputMode outputMode

	inputMode inputMode
	inputs    []input
}

// construct default color state
func newColor() color {
	c := color{
		outputMode: hexOutputMode,

		inputMode: rgbInputMode,
		inputs:    []input{&rgb{values: [3]int{20, 80, 80}}, &hsl{values: [3]int{200, 50, 50}}},
	}
	return c
}

// parse any color format and set inputs accordingly
func (self *color) parseInput(s string) {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	switch {
	case strings.HasPrefix(s, "rgb("):
		rgb, err := rgbFromString(s)
		if err != nil {
			log.Println(err)
		}
		self.inputMode = rgbInputMode
		self.inputs[rgbInputMode] = &rgb
	case strings.HasPrefix(s, "hsl("):
		hsl, err := hslFromString(s)
		if err != nil {
			log.Println(err)
		}
		self.inputMode = hslInputMode
		self.inputs[hslInputMode] = &hsl
	default:
		rgb, err := rgbFromHexString(s)
		if err != nil {
			log.Println(err)
		}
		self.inputMode = rgbInputMode
		self.inputs[rgbInputMode] = &rgb
	}
}

func (self *color) CurrentInput() *input {
	return &self.inputs[self.inputMode]
}

func (self *color) AsRgb() (rgb rgb) {
	return (*self.CurrentInput()).ToRgb()
}

func (self *color) setInputMode(newInputMode inputMode) {
	switch newInputMode {
	case rgbInputMode:
		rgb := self.AsRgb()
		self.inputs[rgbInputMode] = &rgb
	case hslInputMode:
		hsl := rgbToHsl(self.AsRgb())
		self.inputs[hslInputMode] = &hsl
	default:
		log.Fatalf("unexpected inputMode %v", newInputMode)
	}
	self.inputMode = newInputMode
}

func (self *color) CycleInputModes() {
	switch self.inputMode {
	case rgbInputMode:
		self.setInputMode(hslInputMode)
	case hslInputMode:
		self.setInputMode(rgbInputMode)
	default:
		log.Fatalf("unexpected inputMode %v", self.inputMode)
	}
}

func (self *color) ScrollCurrentValue(n int) {
	(*self.CurrentInput()).ScrollCurrentValue(n)
}

func (self *color) ScrollValueIndex(n int) {
	(*self.CurrentInput()).ScrollValueIndex(n)
}

func (self *color) CycleOutputModes() {
	switch self.outputMode {
	case hexOutputMode:
		self.outputMode = rgbOutputMode
	case rgbOutputMode:
		self.outputMode = hslOutputMode
	case hslOutputMode:
		self.outputMode = hexOutputMode
	default:
		log.Panicf("unexpected outputMode %v", self.outputMode)
	}
}

func (self *color) Output() string {
	rgb := (*self.CurrentInput()).ToRgb()
	switch self.outputMode {
	case hexOutputMode:
		return rgbToHexString(rgb)
	case rgbOutputMode:
		return rgbToString(rgb)
	case hslOutputMode:
		return hslToString(rgbToHsl(rgb))
	default:
		log.Panicf("unexpected outputMode %v", self.outputMode)
	}
	return ""
}
