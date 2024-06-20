package app

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

func (self *app) drawText(x int, y int, text string, style tcell.Style) {
	for _, c := range text {
		self.screen.SetContent(x, y, c, nil, style)
		x++
	}
}

func (self *app) drawOutput(x int, y int) {
	const outputTxt = "output: "
	self.drawText(x, y, outputTxt, tcell.StyleDefault)
	x += len(outputTxt)

	output := self.color.output()

	r, g, b := self.color.currentAsRgb().triple()
	style := tcell.StyleDefault.Background(tcell.NewRGBColor(int32(r), int32(g), int32(b))).Foreground(tcell.ColorBlack)

	self.drawText(x, y, output, style)
}

const sliderLen = 64

type slider struct {
	min     int
	max     int
	current int
	// length int
	// step int
	prefix string
}

func tcellColor(r int, g int, b int) tcell.Color {
	return tcell.NewRGBColor(int32(r), int32(g), int32(b))
}

func (self *app) drawSlider(x, y int, sl slider, getRgbAtIndex func(i int) (int, int, int), current bool) {
	style := tcell.StyleDefault
	if !current {
    	style = style.Foreground(tcell.ColorGray) 
	}
	self.drawText(x, y, sl.prefix, style)
	x += len(sl.prefix)

	for i := 0; i < sliderLen; i++ {
		var st tcell.Style
		if i != sl.current*sliderLen/(sl.max-sl.min+1) {
			r, g, b := getRgbAtIndex(i)
			st = tcell.StyleDefault.Background(tcellColor(r, g, b))
		} else {
			st = tcell.StyleDefault.Background(tcell.ColorGray)
		}
		self.screen.SetContent(x+i, y, ' ', nil, st)
	}
}

func (self *app) drawSliders(x, y int) {
	var sl slider

	switch self.color.inputMode {
	case rgbInputMode:
		r, g, b := self.color.rgb.triple()

		sl = slider{min: 0, max: 255, prefix: fmt.Sprintf("R: %03d ", r), current: r}

		getRgbAtIndex := func(i int) (int, int, int) {
			return i * (sl.max - sl.min + 1) / sliderLen, g, b
		}
		self.drawSlider(x, y, sl, getRgbAtIndex, self.color.currentSlider == 0)

		sl = slider{min: 0, max: 255, prefix: fmt.Sprintf("G: %03d ", g), current: g}
		getRgbAtIndex = func(i int) (int, int, int) {
			return r, i * (sl.max - sl.min + 1) / sliderLen, b
		}
		self.drawSlider(x, y+1, sl, getRgbAtIndex, self.color.currentSlider == 1)

		sl = slider{min: 0, max: 255, prefix: fmt.Sprintf("B: %03d ", b), current: b}
		getRgbAtIndex = func(i int) (int, int, int) {
			return r, g, i * (sl.max - sl.min + 1) / sliderLen
		}
		self.drawSlider(x, y+2, sl, getRgbAtIndex, self.color.currentSlider == 2)

	case hslInputMode:
		h, s, l := self.color.hsl.triple()

		sl = slider{min: 0, max: 360, prefix: fmt.Sprintf("H: %03d ", h), current: h}
		getRgbAtIndex := func(i int) (int, int, int) {
			return hslToRgb(hsl{i * (sl.max - sl.min + 1) / sliderLen, s, l}).triple()
		}
		self.drawSlider(x, y, sl, getRgbAtIndex, self.color.currentSlider == 0)

		sl = slider{min: 0, max: 100, prefix: fmt.Sprintf("S: %03d ", s), current: s}
		getRgbAtIndex = func(i int) (int, int, int) {
			return hslToRgb(hsl{h, i * (sl.max - sl.min + 1) / sliderLen, l}).triple()
		}
		self.drawSlider(x, y+1, sl, getRgbAtIndex, self.color.currentSlider == 1)

		sl = slider{min: 0, max: 100, prefix: fmt.Sprintf("L: %03d ", l), current: l}
		getRgbAtIndex = func(i int) (int, int, int) {
			return hslToRgb(hsl{h, s, i * (sl.max - sl.min + 1) / sliderLen}).triple()
		}
		self.drawSlider(x, y+2, sl, getRgbAtIndex, self.color.currentSlider == 2)

	default:
		log.Panicf("unexpected inputMode %v", self.color.inputMode)
	}
}

func (self *app) draw() {
	x, y := 0, 0

	self.drawOutput(x, y)
	y += 1
	self.drawSliders(x, y)
}
