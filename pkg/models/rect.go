package models

type Rect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func NewRect(x, y, width, height int32) *Rect {
	return &Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}
