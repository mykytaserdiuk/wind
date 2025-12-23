package base

import rl "github.com/gen2brain/raylib-go/raylib"

type ColorViewer struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func NewColorViewer() *ColorViewer {
	return &ColorViewer{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
}

func (cp *ColorViewer) SetColor(r, g, b, a uint8) {
	cp.R = r
	cp.G = g
	cp.B = b
	cp.A = a
}

func (cp *ColorViewer) GetColor() rl.Color {
	return rl.Color{
		R: cp.R,
		G: cp.G,
		B: cp.B,
		A: cp.A,
	}
}
