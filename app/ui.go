package app

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func (self *app) shiftSlider(n int64) {
	switch self.currentSlider {
	case 0:
		self.rgb.r = clamp(self.rgb.r+n, 0, 255)
	case 1:
		self.rgb.g = clamp(self.rgb.g+n, 0, 255)
	case 2:
		self.rgb.b = clamp(self.rgb.b+n, 0, 255)
	}
}

func (self *app) downSlider() {
	self.currentSlider = min(2, self.currentSlider+1)
}

func (self *app) upSlider() {
	self.currentSlider = max(0, self.currentSlider-1)
}

func (self *app) handleEvent(ev tcell.Event) (quit bool) {
	switch ev := ev.(type) {
	case *tcell.EventResize:
		self.termX, self.termY = ev.Size()
		self.screen.Sync()
	case *tcell.EventKey:
		switch {
		case ev.Rune() == 'q' || ev.Key() == tcell.KeyCtrlC:
			quit = true
		case ev.Key() == tcell.KeyEnter:
			self.printOnExit = true
			quit = true
		case ev.Rune() == 'h':
			self.shiftSlider(-1)
		case ev.Rune() == 'l':
			self.shiftSlider(+1)
		case ev.Rune() == 'b':
			self.shiftSlider(-8)
		case ev.Rune() == 'w':
			self.shiftSlider(+8)
		case ev.Rune() == '[':
			self.shiftSlider(-256)
		case ev.Rune() == ']':
			self.shiftSlider(+256)
		case ev.Rune() == 'j':
			self.downSlider()
		case ev.Rune() == 'k':
			self.upSlider()
		case ev.Key() == tcell.KeyCtrlL:
			self.screen.Sync()
		}

	case *tcell.EventMouse:
		switch ev.Buttons() {
		case tcell.Button1, tcell.Button2:
		case tcell.ButtonNone:
		}
	}
	return
}

func (self *app) drawText(x int, y int, text string) {
	style := tcell.StyleDefault
	for _, c := range text {
		self.screen.SetContent(x, y, c, nil, style)
		x++
	}
}

func (self *app) drawSelectedColor(x int, y int) {
	colorHex := rgbToHex(self.rgb)

	style := tcell.StyleDefault.Background(tcell.NewRGBColor(int32(self.rgb.r), int32(self.rgb.g), int32(self.rgb.b))).Foreground(tcell.ColorBlack)

	const currentColorText = "current: "
	self.drawText(x, y, currentColorText)
	x += len(currentColorText)

	for _, c := range colorHex {
		self.screen.SetContent(x, y, c, nil, style)
		x += 1
	}
}

func (self *app) drawSlider(x, y int, currentValue int, prefix string, getStyle func(int) tcell.Style) {
	self.drawText(x, y, prefix)

	offset := x+len(prefix)
	for i := 0; i <= 64; i++ {
		if i == currentValue {
			self.screen.SetContent(offset+i, y, '+', nil, tcell.StyleDefault)
		} else {
			self.screen.SetContent(offset+i, y, ' ', nil, getStyle(i))
		}
	}
}

func (self *app) drawSliders(x, y int) {
	self.drawSlider(x, y, int(self.rgb.r/4), "R: "+fmt.Sprintf("%03d", self.rgb.r)+" ", func(i int) tcell.Style {
		return tcell.StyleDefault.Background(tcell.NewRGBColor(int32(i*4), int32(self.rgb.g), int32(self.rgb.b)))
	})
	self.drawSlider(x, y+1, int(self.rgb.g/4), "G: "+fmt.Sprintf("%03d", self.rgb.g)+" ", func(i int) tcell.Style {
		return tcell.StyleDefault.Background(tcell.NewRGBColor(int32(self.rgb.r),  int32(i*4), int32(self.rgb.b)))
	})
	self.drawSlider(x, y+2, int(self.rgb.b/4), "G: "+fmt.Sprintf("%03d", self.rgb.b)+" ", func(i int) tcell.Style {
		return tcell.StyleDefault.Background(tcell.NewRGBColor(int32(self.rgb.r), int32(self.rgb.g), int32(i*4)))
	})

	// self.drawSlider(x, y, int(self.rgb.r/4), "R: "+fmt.Sprintf("%03d", self.rgb.r)+" ", func(i int, displayedLength int) tcell.Style {
	// 	return tcell.StyleDefault.Background(tcell.NewRGBColor(int32(i * 256/displayedLength), 0, 0))
	// })
}

func (self *app) draw() {
	x, y := 0, 0

	self.drawSelectedColor(x, y)
	y += 1
	self.drawSliders(x, y)
}
