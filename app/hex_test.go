package main

import (
	"testing"
)

func TestHexToRgb(t *testing.T) {
	expect := "#00CED1"
	got := rgbToHex(rgb{0, 206, 209})

	if got != expect {
		t.Errorf("got %s, expect %s", got, expect)
	}
}

func TestRgbToHex(t *testing.T) {
	expect := rgb{0, 206, 209}
	got, err := hexToRgb("#00CED1")

	if err != nil {
		t.Fatal(err)
	}
	if got != expect {
		t.Errorf("got %v, expect %v", got, expect)
	}
}
