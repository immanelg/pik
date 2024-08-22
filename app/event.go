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
        key, r := ev.Key(), ev.Rune()
		switch {
		case r == 'q' || key == tcell.KeyCtrlC:
			quit = true
		case key == tcell.KeyEnter:
			self.printOnExit = true
			quit = true

		case key == tcell.KeyCtrlL:
			self.screen.Sync()

		case r == 'h':
			self.color.ScrollCurrentValue(-1)
		case r == 'l':
			self.color.ScrollCurrentValue(+1)
		case r == 'b':
			self.color.ScrollCurrentValue(-8)
		case r == 'w':
			self.color.ScrollCurrentValue(+8)
		case r == '[':
			self.color.ScrollCurrentValue(-32)
		case r == ']':
			self.color.ScrollCurrentValue(+32)
		case r == 'H':
			self.color.ScrollCurrentValueToBound(false)
		case r == 'L':
			self.color.ScrollCurrentValueToBound(true)

		case r == 'j':
			self.color.ScrollValueIndex(+1)
		case r == 'k':
			self.color.ScrollValueIndex(-1)

		case r == 'i':
			self.color.CycleInputModes()

		case r == 'o':
			self.color.CycleOutputModes()

		case r == 'y':
			c := self.color.Output()
			go func() {
				if err := clipboard.Set(c); err != nil {
					log.Printf("error writing clipboard: %v", err)
				}
			}()

		case r == 'p':
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
