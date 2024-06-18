package app

import (
	"log"
)

// NOTE: hsl<->rgb<->whatever is not a bijection.
// We can't have only one piece of state for color and then have different "views" to them (rgb/hsl/hsv/whatever).
// When changing a hsl slider value in this model, we have to convert rgb->hsl on the fly, change hsl, and then
// convert it back; it will lose information.

// Displayed color mode
type inputMode uint8

const (
	rgbInputMode = inputMode(iota)
	hslInputMode
)

// What will be printed after exiting
type outputMode uint8

const (
	hexOutputMode = outputMode(iota)
	rgbOutputMode
	hslOutputMode
)

// Manages the current color state actions.
// Input modes is synchronized on mode change.
type color struct {
	currentSlider int

	outputMode outputMode

	inputMode inputMode
	rgb       rgb
	hsl       hsl
}

func (self *color) currentAsRgb() (rgb rgb) {
	switch self.inputMode {
	case rgbInputMode:
		rgb = self.rgb
	case hslInputMode:
		rgb = hslToRgb(self.hsl)
	default:
		log.Panicf("unexpected inputMode %v", self.inputMode)
	}
	return
}

func (self *color) cycleInputMode() {
	switch self.inputMode {
	case rgbInputMode:
		self.setInputMode(hslInputMode)
	case hslInputMode:
		self.setInputMode(rgbInputMode)
	default:
		log.Panicf("unexpected inputMode %v", self.inputMode)
	}
}

func (self *color) setInputMode(newInputMode inputMode) {
	// Update the state of the new mode
	// calculating it from the old one.
	switch newInputMode {
	case rgbInputMode:
		self.rgb = self.currentAsRgb()
	case hslInputMode:
		self.hsl = rgbToHsl(self.currentAsRgb())
	default:
		log.Panicf("unexpected inputMode %v", newInputMode)
	}
	self.inputMode = newInputMode
	self.currentSlider = 0
}

func (self *color) changeSliderValue(n int) {
	// scroll by percent?
	// if !(perc >= -1 && perc <= 1) {
	// 	log.Panicf("perc is not in range [-1,1]: %v", perc)
	// }
	switch self.inputMode {
	case rgbInputMode:
		switch self.currentSlider {
		case 0:
			self.rgb.r = clamp(self.rgb.r+n, 0, 255)
		case 1:
			self.rgb.g = clamp(self.rgb.g+n, 0, 255)
		case 2:
			self.rgb.b = clamp(self.rgb.b+n, 0, 255)
		default:
			log.Panicf("unexpected currentSlider %v", self.currentSlider)
		}
	case hslInputMode:
		switch self.currentSlider {
		case 0:
			self.hsl.h = clamp(self.hsl.h+n, 0, 360)
		case 1:
			self.hsl.s = clamp(self.hsl.s+n, 0, 100)
		case 2:
			self.hsl.l = clamp(self.hsl.l+n, 0, 100)
		default:
			log.Panicf("unexpected currentSlider %v", self.currentSlider)
		}
	default:
		log.Panicf("unexpected mode %v", self.inputMode)
	}
}

func (self *color) changeCurrentSlider(n int) {
	self.currentSlider = clamp(self.currentSlider+n, 0, 2)
}

func (self *color) cycleOutputMode() {
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

func (self *color) getOutput() string {
	rgb := self.currentAsRgb()

	switch self.outputMode {
	case hexOutputMode:
		return rgbToHex(rgb)
	case rgbOutputMode:
		return "rgb(...)"
		// return rgbString() // TODO
	case hslOutputMode:
		return "hsl(...)"
		// return hslString() // TODO
	default:
		log.Panicf("unexpected outputMode %v", self.outputMode)
	}
	return ""
}
