package app

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

const sliderLen = 32

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

	output := self.color.Output()

	r, g, b := self.color.AsRgb().Triple()
	style := tcell.StyleDefault.Background(tcell.NewRGBColor(int32(r), int32(g), int32(b))).Foreground(tcell.ColorBlack)

	self.drawText(x, y, output, style)
}

func tcellColor(r int, g int, b int) tcell.Color {
	return tcell.NewRGBColor(int32(r), int32(g), int32(b))
}

func (self *app) drawSliders(x, y int) {

	inp := *self.color.CurrentInput()
	prefix := inp.Prefix()
	min := inp.Min()
	max := inp.Max()

	for i, value := range inp.Values() {
		style := tcell.StyleDefault.Foreground(tcell.ColorGray)
		if i == inp.CurrentValueIndex() {
			style = tcell.StyleDefault.Bold(true)
		}
		txt := fmt.Sprintf("%s: %-3d ", prefix[i], value)
		self.drawText(x, y+i, txt, style)
		xOffset := len(txt)

		for j := 0; j < sliderLen; j++ {
			var st tcell.Style
			if j != value*sliderLen/(max[i]-min[i]+1) {
				c := tcellColor(inp.WithValue(i, j*(max[i]-min[i]+1)/sliderLen).ToRgb().Triple())
				st = tcell.StyleDefault.Background(c)
			} else {
				st = tcell.StyleDefault.Background(tcell.ColorGray)
			}
			self.screen.SetContent(x+xOffset+j, y+i, ' ', nil, st)
		}
	}
}

func (self *app) draw() {
	x, y := 0, 0
	self.drawOutput(x, y)
	self.drawSliders(x, y+1)
}
