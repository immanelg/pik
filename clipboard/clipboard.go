package clipboard

import (
	"errors"
	"io"
	"os"
	"os/exec"
)

var PasteSupported = false
var CopySupported = false

var copyCmd []string
var pasteCmd []string

var errUnsupported = errors.New("unsupported clipboard")

func Guess() {
	switch {
	case os.Getenv("WAYLAND_DISPLAY") != "":
		if path, err := exec.LookPath("wl-copy"); err == nil {
			CopySupported = true
			copyCmd = []string{path}
		}
		if path, err := exec.LookPath("wl-paste"); err == nil {
			PasteSupported = true
			pasteCmd = []string{path, "-n"}
		}

	case os.Getenv("DISPLAY") != "":
		if path, err := exec.LookPath("xclip"); err == nil {
			PasteSupported = true
			CopySupported = true
			copyCmd = []string{path, "-selection", "clipboard"}
			pasteCmd = []string{path, "-selection", "clipboard", "-o"}
		}
	default:
	}
}

func Set(data string) error {
	if !CopySupported {
		return errUnsupported
	}
	cmd := exec.Command(copyCmd[0], copyCmd[1:]...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	defer stdin.Close()

	_, err = io.WriteString(stdin, data)
	if err != nil {
	    return err
	} 
	_, err = cmd.CombinedOutput()
	return err
}

func Get() (string, error) {
	if !PasteSupported {
		return "", errUnsupported

	}
	cmd := exec.Command(pasteCmd[0], pasteCmd[1:]...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
