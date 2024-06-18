package app

type rgb struct {
	r int // 0..255
	g int // 0..255
	b int // 0.255
}

type rgba struct {
	rgb
	a int // 0..100
}

func (self rgb) inverted() rgb {
	return rgb{
		r: 255 - self.r,
		g: 255 - self.g,
		b: 255 - self.b,
	}
}

func (self rgb) triple() (int, int, int) {
	return self.r, self.g, self.b
}

var whiteRgb = rgb{255, 255, 255}
var blackRgb = rgb{0, 0, 0}
var initRgb = rgb{216, 146, 215}
