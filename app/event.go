package app

import (
	"github.com/gdamore/tcell/v2"
)

func (self *app) handleEvent(ev tcell.Event) (quit bool) {
	switch ev := ev.(type) {
	case *tcell.EventResize:
		self.termW, self.termH = ev.Size()
		self.screen.Sync()
	case *tcell.EventKey:
		switch {
		case ev.Rune() == 'q' || ev.Key() == tcell.KeyCtrlC:
			quit = true
		case ev.Key() == tcell.KeyEnter:
			self.printOnExit = true
			quit = true

		case ev.Rune() == 'h':
			self.color.changeSliderValue(-1)
		case ev.Rune() == 'l':
			self.color.changeSliderValue(+1)
		case ev.Rune() == 'b':
			self.color.changeSliderValue(-8)
		case ev.Rune() == 'w':
			self.color.changeSliderValue(+8)
		case ev.Rune() == '[':
			self.color.changeSliderValue(-32)
		case ev.Rune() == ']':
			self.color.changeSliderValue(+32)

		case ev.Rune() == 'j':
			self.color.changeCurrentSlider(+1)
		case ev.Rune() == 'k':
			self.color.changeCurrentSlider(-1)

		case ev.Rune() == 'i':
			self.color.cycleInputMode()
		case ev.Rune() == 'o':
			self.color.cycleOutputMode()

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
