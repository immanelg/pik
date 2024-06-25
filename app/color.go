package app

import (
	"log"
)

// manages state of input and output
type color struct {
	outputMode outputMode

	// current index for inputs
	inputMode inputMode
	// color inputs, should have length of set of inputMode enum values
	inputs [inputModeCount]input
}

// construct default color state.  you need to call ParseInput or Set<...> to change defaults
func newColor() color {
	defaulthsl := newHsl(100, 50, 50)

	c := color{}
	c.inputMode = hslInputMode
	c.inputs[rgbInputMode] = &rgb{}
	c.inputs[hslInputMode] = &defaulthsl

	c.outputMode = hexOutputMode
	return c
}

func (self *color) ParseInput(s string) {
	inputMode, input, err := parseInput(s)
	if err != nil {
		log.Println(err)
	}
	self.SetInput(inputMode, input)
}

func (self *color) SetInput(inputMode inputMode, input input) {
	self.inputMode = inputMode
	self.inputs[inputMode] = input
}

func (self *color) SetRgb(rgb rgb) {
	self.inputMode = rgbInputMode
	self.inputs[rgbInputMode] = &rgb
}

func (self *color) SetHsl(hsl hsl) {
	self.inputMode = hslInputMode
	self.inputs[rgbInputMode] = &hsl
}

// get current input
func (self *color) CurrentInput() *input {
	return &self.inputs[self.inputMode]
}

// convert current color to rgb
func (self *color) AsRgb() (rgb rgb) {
	return (*self.CurrentInput()).ToRgb()
}

// change the color mode, converting from current mode
func (self *color) SwitchInputMode(newInputMode inputMode) {
	switch newInputMode {
	case rgbInputMode:
		rgb := self.AsRgb()
		self.SetRgb(rgb)
	case hslInputMode:
		hsl := rgbToHsl(self.AsRgb())
		self.SetHsl(hsl)
	default:
		log.Fatalf("unexpected inputMode %v", newInputMode)
	}
}

func (self *color) CycleInputModes() {
	switch self.inputMode {
	case rgbInputMode:
		self.SwitchInputMode(hslInputMode)
	case hslInputMode:
		self.SwitchInputMode(rgbInputMode)
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

// serialize color to pretty string
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
