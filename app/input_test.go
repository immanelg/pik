package app

import "testing"

func TestParseInputGood(t *testing.T) {
	makeRgb := func(r int, g int, b int) input {
		i := newRgb(r, g, b)
		return &i
	}
	makeHsl := func(h int, s int, l int) input {
		i := newHsl(h, s, l)
		return &i
	}

	testCases := []struct {
		inp             string
		expectInputMode inputMode
		expectInput     input
	}{
		{"rgb(100 200 255)", rgbInputMode, makeRgb(100, 200, 255)},
		{"#00FF00", rgbInputMode, makeRgb(0, 255, 0)},
		{"hsl(300 100 100)", hslInputMode, makeHsl(300, 100, 100)},
	}

	for _, tc := range testCases {
		inputMode, input, err := parseInput(tc.inp)
		if err != nil {
			t.Errorf("input %+v, got error: %+v", tc.inp, err)
			continue
		}

		if inputMode != tc.expectInputMode {
			t.Errorf("input %+v, expected inputMode: %+v, got: %+v", tc.inp, tc.expectInputMode, inputMode)
			continue
		}

		if input.Values() != tc.expectInput.Values() {
			t.Errorf("input %v, expected input.Values(): %+v, got: %+v", tc.inp, tc.expectInput.Values(), input.Values())
			continue
		}

	}
}
func TestParseInputBad(t *testing.T) {
	testCases := []struct {
		inp string
	}{
		{"nonsense"},
		{"rgb(1000 2 3)"},
		{"hsl(100 19000 299)"},
		{"hsl(100 100 100 100)"},
		{"1234567890"},
	}

	for _, tc := range testCases {
		inputMode, input, err := parseInput(tc.inp)
		if err == nil {
			t.Errorf("input %v, expected error, got inputMode %v, input %v", tc.inp, inputMode, input)
			continue
		}

	}
}
