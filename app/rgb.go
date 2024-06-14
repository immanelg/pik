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

var white = rgb{255, 255, 255}

type rgba struct {
	rgb
	a int64
}
