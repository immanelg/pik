package app

type rgb struct {
	r, g, b int64
}

func (self rgb) inverted() rgb {
	return rgb{
		r: 255 - self.r,
		g: 255 - self.g,
		b: 255 - self.b,
	}
}

var whiteRgb = rgb{255, 255, 255}
var blackRgb = rgb{0, 0, 0}

type rgba struct {
	rgb
	a int64
}
