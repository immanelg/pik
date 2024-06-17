package app

import (
	"testing"
)

func TestHexToRgb(t *testing.T) {
	testCases := []struct {
		input  string
		expect rgb
	}{
		{"#00CED1", rgb{0, 206, 209}},
		{"#348217", rgb{52, 130, 23}},
	}

	for _, tc := range testCases {
		result, err := hexToRgb(tc.input)
		if result != tc.expect {
			t.Errorf("input %v, expect: %v, got: %v", tc.input, tc.expect, result)
		}
		if err != nil {
			t.Errorf("input %v, got err: %v", tc.input, err)
		}
	}

}

func TestRgbToHex(t *testing.T) {
	testCases := []struct {
		input  rgb
		expect string
	}{
		{rgb{0, 206, 209}, "#00CED1"},
		{rgb{52, 130, 23}, "#348217"},
	}

	for _, tc := range testCases {
		result := rgbToHex(tc.input)
		if result != tc.expect {
			t.Errorf("input %v, expect: %v, got: %v", tc.input, tc.expect, result)
		}
	}
}
