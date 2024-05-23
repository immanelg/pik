package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type Mode int8

const (
	RgbMode Mode = 0
	HslMode Mode = 1
)

// application state
type app struct {
	screen        tcell.Screen
	termX, termY  int
	rgb           rgb
	mode          Mode
	currentSlider int
	printOnExit bool
}

func newApp(initialColor rgb) app {
	return app{rgb: initialColor}
}

// inits terminal ui and runs event loop.
func (self *app) loop() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	self.screen = screen
	self.termX, self.termY = self.screen.Size()

	// do not leave terminal in a weird state
	cleanup := func() {
		err := recover()
		screen.Fini()
		if err != nil {
			panic(err)
		}
	}
	defer cleanup()

	defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	self.screen.SetStyle(defaultStyle)
	self.screen.EnableMouse()
	self.screen.Clear()

	for {
		self.draw()
		self.screen.Show()

		ev := self.screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			self.termX, self.termY = ev.Size()
			self.screen.Sync()
		case *tcell.EventKey:
			if ev.Rune() == 'q' || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Rune() == 'h' {
				switch self.currentSlider {
				case 0:
					self.rgb.r = max(0, self.rgb.r-1)
				case 1:
					self.rgb.g = max(0, self.rgb.g-1)
				case 2:
					self.rgb.b = max(0, self.rgb.b-1)
				}
			} else if ev.Rune() == 'l' {
				switch self.currentSlider {
				case 0:
					self.rgb.r = min(255, self.rgb.r+1)
				case 1:
					self.rgb.g = min(255, self.rgb.g+1)
				case 2:
					self.rgb.b = min(255, self.rgb.b+1)
				}
			} else if ev.Rune() == 'j' {
				self.currentSlider = min(2, self.currentSlider+1)
			} else if ev.Rune() == 'k' {
				self.currentSlider = max(0, self.currentSlider-1)
			} else if ev.Key() == tcell.KeyCtrlL {
				self.screen.Sync()
			} else if ev.Key() == tcell.KeyEnter {
				self.printOnExit = true
				return
			}

		case *tcell.EventMouse:
			switch ev.Buttons() {
			case tcell.Button1, tcell.Button2:
			case tcell.ButtonNone:
			}
		}
	}
}

func (self *app) drawSelectedColor(x int, y int) {
	colorHex := rgbToHex(self.rgb)

	inverted := self.rgb.inverted()
	style := tcell.StyleDefault.Background(
		tcell.NewRGBColor(int32(self.rgb.r), int32(self.rgb.g), int32(self.rgb.b))).Foreground(
		tcell.NewRGBColor(int32(inverted.r), int32(inverted.g), int32(inverted.b)))

	for {
		if x > 9 {
			break
		}
		self.screen.SetContent(x, y, ' ', nil, style)
		x += 1
	}

	self.screen.SetContent(x, y, ' ', nil, tcell.StyleDefault)
	for _, c := range colorHex {
		self.screen.SetContent(x, y, c, nil, tcell.StyleDefault)
		x += 1
	}
}

func (self *app) drawSliders(x, y int) {
	for x := x; x <= 64; x++ {
		xScaled := int32(x * 4)
		style := tcell.StyleDefault
		self.screen.SetContent(x, y, ' ', nil, style.Background(tcell.NewRGBColor(xScaled, 0, 0)))
		self.screen.SetContent(x, y+1, ' ', nil, style.Background(tcell.NewRGBColor(0, xScaled, 0)))
		self.screen.SetContent(x, y+2, ' ', nil, style.Background(tcell.NewRGBColor(0, 0, xScaled)))
	}

	redCursor := int32(self.rgb.r / 4)
	greenCursor := int32(self.rgb.g / 4)
	blueCursor := int32(self.rgb.b / 4)
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorDefault)
	self.screen.SetContent(x+int(redCursor), y, 'X', nil, style.Foreground(tcell.NewRGBColor(redCursor, 0, 0)))
	self.screen.SetContent(x+int(greenCursor), y+1, 'X', nil, style.Foreground(tcell.NewRGBColor(0, greenCursor, 0)))
	self.screen.SetContent(x+int(blueCursor), y+2, 'X', nil, style.Foreground(tcell.NewRGBColor(0, 0, blueCursor)))
}

func (self *app) draw() {
	x, y := 0, 0

	self.drawSelectedColor(x, y)
	y += 1
	self.drawSliders(x, y)
}
