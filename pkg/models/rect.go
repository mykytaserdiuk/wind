package models

import rl "github.com/gen2brain/raylib-go/raylib"

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

func (r *Rect) GetBounds() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(r.PosX),
		Y:      float32(r.PosY),
		Width:  float32(r.Width),
		Height: float32(r.Height),
	}
}
