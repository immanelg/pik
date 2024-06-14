package app

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
	printOnExit   bool
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
		quit := self.handleEvent(ev)
		if quit {
			break
		}
	}
}
