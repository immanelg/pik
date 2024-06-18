package app

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func Run() {
	var initialColor string
	var logfile string
	flag.StringVar(&initialColor, "hex", "", "initial color to edit (hex)")
	flag.StringVar(&logfile, "logfile", "", "file for debug logging")
	flag.Parse()

	var initialRgb rgb = initRgb
	if initialColor != "" {
		var err error
		initialRgb, err = hexToRgb(initialColor)
		if err != nil {
			log.Fatalf("cannot parse HEX color: %s\n", err)
		}
	}

	if logfile != "" {
		f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println("started")
	} else {
		log.SetOutput(io.Discard)
	}

	app := newApp(initialRgb)
	app.loop()

	if app.printOnExit {
		fmt.Println(app.color.output())
	}
}
