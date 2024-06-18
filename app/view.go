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
	const currentColorText = "output: "
	self.drawText(x, y, currentColorText, tcell.StyleDefault)
	x += len(currentColorText)

	output := self.color.getOutput()

	r, g, b := self.color.rgb.triple()
	style := tcell.StyleDefault.Background(tcell.NewRGBColor(int32(r), int32(g), int32(b))).Foreground(tcell.ColorBlack)

	self.drawText(x, y, output, style)
}

func (self *app) drawSlider(x, y int, selectedIndex int, prefix string, getStyle func(int) tcell.Style) {
	style := tcell.StyleDefault
	self.drawText(x, y, prefix, style)

	offset := x + len(prefix)
	for i := 0; i < 64; i++ {
		if i == selectedIndex {
			self.screen.SetContent(offset+i, y, '*', nil, tcell.StyleDefault.Background(tcell.ColorGray))
		} else {
			self.screen.SetContent(offset+i, y, ' ', nil, getStyle(i))
		}
	}
}

func (self *app) drawSliders(x, y int) {
	switch self.color.inputMode {
	case rgbInputMode:
		r, g, b := self.color.rgb.triple()
		sliderMax := 256
		sliderUiLen := 64
		self.drawSlider(x, y, int(r*sliderUiLen/sliderMax), fmt.Sprintf("R: %03d ", r), func(i int) tcell.Style {
			return tcell.StyleDefault.Background(tcell.NewRGBColor(int32(i*sliderMax/sliderUiLen), int32(g), int32(b)))
		})
		self.drawSlider(x, y+1, int(g*sliderUiLen/sliderMax), fmt.Sprintf("G: %03d ", g), func(i int) tcell.Style {
			return tcell.StyleDefault.Background(tcell.NewRGBColor(int32(r), int32(i*sliderMax/sliderUiLen), int32(b)))
		})
		self.drawSlider(x, y+2, int(b*sliderUiLen/sliderMax), fmt.Sprintf("B: %03d ", b), func(i int) tcell.Style {
			return tcell.StyleDefault.Background(tcell.NewRGBColor(int32(r), int32(g), int32(i*sliderMax/sliderUiLen)))
		})

	case hslInputMode:
		h, s, l := self.color.hsl.triple()

		self.drawSlider(x, y, int(h*64/360), fmt.Sprintf("H: %03d ", h), func(i int) tcell.Style {
			h, s, l := hslToRgb(hsl{i * 360 / 64, s, l}).triple()
			style := tcell.NewRGBColor(int32(h), int32(s), int32(l))
			return tcell.StyleDefault.Background(style)
		})
		self.drawSlider(x, y+1, int(s*64/100), fmt.Sprintf("S: %03d ", s), func(i int) tcell.Style {
			h, s, l := hslToRgb(hsl{h, i * 100 / 64, l}).triple()
			style := tcell.NewRGBColor(int32(h), int32(s), int32(l))
			return tcell.StyleDefault.Background(style)
		})
		self.drawSlider(x, y+2, int(l*64/100), fmt.Sprintf("L: %03d ", l), func(i int) tcell.Style {
			h, s, l := hslToRgb(hsl{h, s, l*100/64}).triple()
			style := tcell.NewRGBColor(int32(h), int32(s), int32(l))
			return tcell.StyleDefault.Background(style)
		})

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
