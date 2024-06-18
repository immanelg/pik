package app

import (
    "github.com/gdamore/tcell/v2"
)

type config struct{
	keybindings map[tcell.Key]string
	length int
	cursorChar rune
}

// Read only, initialized on start
var cfg config
