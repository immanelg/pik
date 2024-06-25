package app

// what will be printed after exiting
type outputMode uint8

const (
	hexOutputMode = outputMode(iota)
	rgbOutputMode
	hslOutputMode
)
