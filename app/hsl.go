package app

type hsl struct {
	h int // 0..360
	s int // 0..100
	l int // 0.100
}

type hsla struct {
	hsl
	a int // 0..100
}

func (self hsl) triple() (int, int, int) {
	return self.h, self.s, self.l
}

func hslToRgb(hsl hsl) rgb {
	h := float64(hsl.h) / 365.0
	s := float64(hsl.s) / 100.0
	l := float64(hsl.l) / 100.0

	var r, g, b float64

	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - (l * s)
		}
		p := 2*l - q

		r = hueToRgb(p, q, h+1.0/3.0)
		g = hueToRgb(p, q, h)
		b = hueToRgb(p, q, h-1.0/3.0)
	}

	return rgb{int(r * 255), int(g * 255), int(b * 255)}
}

func rgbToHsl(rgb rgb) hsl {
	r := float64(rgb.r) / 255
	g := float64(rgb.g) / 255
	b := float64(rgb.b) / 255

	maxVal := max(max(r, g), b)
	minVal := min(min(r, g), b)

	var h, s, l float64

	l = (maxVal + minVal) / 2

	if maxVal == minVal {
		h = 0
		s = 0
	} else {
		d := maxVal - minVal

		if l > 0.5 {
			s = d / (2 - maxVal - minVal)
		} else {
			s = d / (maxVal + minVal)
		}

		switch maxVal {
		case r:
			h = (g - b) / d
			if g < b {
				h += 6
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}

		h *= 60
	}

	return hsl{int(h), int(s * 100), int(l * 100)}
}

func hueToRgb(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*6*(2.0/3.0-t)
	}
	return p
}
