package app

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

// application state
type app struct {
	screen       tcell.Screen
	termW, termH int
	color        color
	printOnExit  bool
}

func newApp(initialColor rgb) app {
	return app{
		color: color{inputMode: rgbInputMode, rgb: initialColor},
	}
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
	self.termW, self.termH = self.screen.Size()

	// restore terminal
	cleanup := func() {
		err := recover()
		screen.Fini()
		if err != nil {
			panic(err)
		}
	}
	defer cleanup()

	defaultStyle := tcell.StyleDefault

	self.screen.SetStyle(defaultStyle)
	self.screen.EnableMouse()
	self.screen.Clear()

	for {
		self.screen.Fill(' ', tcell.StyleDefault)
		self.draw()
		self.screen.Show()

		ev := self.screen.PollEvent()
		quit := self.handleEvent(ev)
		if quit {
			break
		}
	}
}
