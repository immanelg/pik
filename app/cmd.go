package app

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/immanelg/pik/clipboard"
)

func Run() {
	var inputColor string
	var logfile string
	flag.StringVar(&inputColor, "init", "", "initial color to edit")
	flag.StringVar(&logfile, "logfile", "", "file for debug logging")
	flag.Parse()

	// if not given -init, get input from stdin
	if inputColor == "" {
		fi, _ := os.Stdin.Stat()

		if (fi.Mode() & os.ModeCharDevice) == 0 {
			r := bufio.NewReader(os.Stdin)
			input, err := r.ReadString('\n')
			if err != nil {
				log.Printf("error while reading stdin: %v", err)
			}
			inputColor = input
		}
	}
	color := color{}
	if inputColor != "" {
		color = colorFromInput(inputColor)
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

	clipboard.Guess()

	app := newApp(color)

	app.loop()

	if app.printOnExit {
		fmt.Println(app.color.output())
	}
}
