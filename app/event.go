package app

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/immanelg/pik/clipboard"
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

		case ev.Key() == tcell.KeyCtrlL:
			self.screen.Sync()

		case ev.Rune() == 'h':
			self.color.ScrollCurrentValue(-1)
		case ev.Rune() == 'l':
			self.color.ScrollCurrentValue(+1)
		case ev.Rune() == 'b':
			self.color.ScrollCurrentValue(-8)
		case ev.Rune() == 'w':
			self.color.ScrollCurrentValue(+8)
		case ev.Rune() == '[':
			self.color.ScrollCurrentValue(-32)
		case ev.Rune() == ']':
			self.color.ScrollCurrentValue(+32)

		case ev.Rune() == 'j':
			self.color.ScrollValueIndex(+1)
		case ev.Rune() == 'k':
			self.color.ScrollValueIndex(-1)

		case ev.Rune() == 'i':
			self.color.CycleInputModes()

		case ev.Rune() == 'o':
			self.color.CycleOutputModes()

		case ev.Rune() == 'y':
			c := self.color.Output()
			go func() {
				if err := clipboard.Set(c); err != nil {
					log.Printf("error writing clipboard: %v", err)
				}
			}()

		case ev.Rune() == 'p':
			if c, err := clipboard.Get(); err != nil {
				log.Printf("error reading clipboard: %v", err)
			} else if c != "" {
				self.color.ParseInput(c)
			}
		}

	case *tcell.EventMouse:
		switch ev.Buttons() {
		case tcell.Button1, tcell.Button2:
		case tcell.ButtonNone:
		}
	}
	return
}
