package models

type Rect struct {
	PosX   int32
	PosY   int32
	Width  int32
	Height int32
}

func NewRect(x, y, width, height int32) *Rect {
	return &Rect{
		PosX:   x,
		PosY:   y,
		Width:  width,
		Height: height,
	}
}

func (r *Rect) Clone() *Rect {
	return &Rect{
		PosX:   r.PosX,
		PosY:   r.PosY,
		Width:  r.Width,
		Height: r.Height,
	}
}
