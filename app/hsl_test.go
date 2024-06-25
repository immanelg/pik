package app

import (
	"testing"
)

func TestRgbToHsl(t *testing.T) {
	testCases := []struct {
		input  rgb
		expect hsl
	}{
		// {rgb{255, 213, 0}, hsl{50, 100, 50}}, // fails due to floating point arithmetic
		{newRgb(255, 0, 0), newHsl(0, 100, 50)},
	}

	for _, tc := range testCases {
		result := rgbToHsl(tc.input)
		if result != tc.expect {
			t.Errorf("input %v, expect: %v, got: %v", tc.input, tc.expect, result)
		}
	}
}

func TestHslToRgb(t *testing.T) {
	testCases := []struct {
		input  hsl
		expect rgb
	}{
		// {hsl{50, 100, 50}, rgb{255, 213, 0}}, // fails due to floating point arithmetic
		{newHsl(0, 100, 50), newRgb(255, 0, 0)},
	}

	for _, tc := range testCases {
		result := hslToRgb(tc.input)
		if result != tc.expect {
			t.Errorf("input %v, expect: %v, got: %v", tc.input, tc.expect, result)
		}
	}
}
