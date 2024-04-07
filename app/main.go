package main

import (
	"flag"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

type Mode int8

const (
	RgbMode Mode = 0
	HslMode Mode = 1
)

// application state
type app struct {
	screen       tcell.Screen
	termX, termY int
	rgb          rgb
	mode         Mode
}

// the main loop of the application.
// receives events and updates UI.
func (self *app) loop() {

	defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	self.screen.SetStyle(defaultStyle)
	self.screen.EnableMouse()
	self.screen.EnablePaste()
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
				self.rgb.r = max(0, self.rgb.r-8)
			} else if ev.Rune() == 'l' {
				self.rgb.r = min(255, self.rgb.r+8)
			} else if ev.Key() == tcell.KeyCtrlL {
				self.screen.Sync()
			}

		case *tcell.EventMouse:
			switch ev.Buttons() {
			case tcell.Button1, tcell.Button2:
			case tcell.ButtonNone:
			}
		}
	}
}

func (self *app) draw() {
	color := rgbToHex(self.rgb)

	inverted := self.rgb.inverted()
	style := tcell.StyleDefault.Background(
		tcell.NewRGBColor(int32(self.rgb.r), int32(self.rgb.g), int32(self.rgb.b))).Foreground(
		tcell.NewRGBColor(int32(inverted.r), int32(inverted.g), int32(inverted.b)))

	x, y := 0, 0
	for _, c := range color {
		self.screen.SetContent(x, y, c, []rune{}, style)
		x += 1
	}

	y = 1
	for x := 0; x <= 64; x++ {
		offset := int32(x*4)
		styleG := tcell.StyleDefault.Background(tcell.NewRGBColor(0, offset, 0))
		styleR := tcell.StyleDefault.Background(tcell.NewRGBColor(offset, 0, 0))
		styleB := tcell.StyleDefault.Background(tcell.NewRGBColor(0, 0, offset))
		self.screen.SetContent(x, y, ' ', []rune{}, styleR)
		self.screen.SetContent(x, y+1, ' ', []rune{}, styleG)
		self.screen.SetContent(x, y+2, ' ', []rune{}, styleB)
	}
}

func main() {
	var initialColor string
	flag.StringVar(&initialColor, "hex", "", "initial color to edit (hex)")
	flag.Parse()

	initialRgb := whiteRgb()
	if initialColor != "" {
		var err error
		initialRgb, err = hexToRgb(initialColor)
		if err != nil {
			log.Fatalf("cannot parse HEX color: %s\n", err)
			os.Exit(1)
		}
	}

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	cleanup := func() {
		err := recover()
		screen.Fini()
		if err != nil {
			panic(err)
		}
	}
	defer cleanup()

	termX, termY := screen.Size()

	app := app{
		screen: screen,
		rgb:    initialRgb,
		termX:  termX,
		termY:  termY,
	}

	app.loop()
}
