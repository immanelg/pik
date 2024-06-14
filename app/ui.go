package app

import (
    "fmt"
	"github.com/gdamore/tcell/v2"
)

func (self *app) shiftSlider(n int64) {
	switch self.currentSlider {
	case 0:
		self.rgb.r = min(255, max(0, self.rgb.r+n))
	case 1:
		self.rgb.g = min(255, max(0, self.rgb.g+n))
	case 2:
		self.rgb.b = min(255, max(0, self.rgb.b+n))
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

func (self *app) drawSliders(x, y int) {
	var rText = "R: " + fmt.Sprintf("%03d", self.rgb.r) + " "
	var gText = "G: " + fmt.Sprintf("%03d", self.rgb.g) + " "
	var bText = "B: " + fmt.Sprintf("%03d", self.rgb.b) + " "
	self.drawText(x, y, rText)
	self.drawText(x, y+1, gText)
	self.drawText(x, y+2, bText)

	for x := x; x <= 64; x++ {
		xScaled := int32(x * 4)
		style := tcell.StyleDefault
		self.screen.SetContent(x+len(rText), y, ' ', nil, style.Background(tcell.NewRGBColor(xScaled, 0, 0)))
		self.screen.SetContent(x+len(gText), y+1, ' ', nil, style.Background(tcell.NewRGBColor(0, xScaled, 0)))
		self.screen.SetContent(x+len(bText), y+2, ' ', nil, style.Background(tcell.NewRGBColor(0, 0, xScaled)))
	}

	redCursor := int32(self.rgb.r / 4)
	greenCursor := int32(self.rgb.g / 4)
	blueCursor := int32(self.rgb.b / 4)
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	self.screen.SetContent(x+len(rText)+int(redCursor), y, 'X', nil, style)
	self.screen.SetContent(x+len(gText)+int(greenCursor), y+1, 'X', nil, style)
	self.screen.SetContent(x+len(bText)+int(blueCursor), y+2, 'X', nil, style)
}

func (self *app) draw() {
	x, y := 0, 0

	self.drawSelectedColor(x, y)
	y += 1
	self.drawSliders(x, y)
}
