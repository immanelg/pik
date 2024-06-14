package app

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func Run() {
	var initialColor string
	flag.StringVar(&initialColor, "hex", "", "initial color to edit (hex)")
	flag.Parse()

	var initialRgb rgb
	if initialColor != "" {
		var err error
		initialRgb, err = hexToRgb(initialColor)
		if err != nil {
			log.Fatalf("cannot parse HEX color: %s\n", err)
			os.Exit(1)
		}
	}

	app := newApp(initialRgb)
	app.loop()

	if app.printOnExit {
		fmt.Println(rgbToHex(app.rgb))
	}
}
