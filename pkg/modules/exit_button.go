package modules

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/pkg/models"
)

type ExitButton struct {
	Rect      *models.Rect
	isDeleted bool
}

func NewExitButton(rect *models.Rect) *ExitButton {
	return &ExitButton{
		Rect: rect,
	}
}

func (eb *ExitButton) GetBounds() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(eb.Rect.PosX),
		Y:      float32(eb.Rect.PosY),
		Width:  float32(eb.Rect.Width),
		Height: float32(eb.Rect.Height),
	}
}

func (eb *ExitButton) Draw() {
	rl.DrawRectangle(eb.Rect.PosX, eb.Rect.PosY, eb.Rect.Width, eb.Rect.Height, rl.Red)
	rl.DrawRectangleLines(eb.Rect.PosX, eb.Rect.PosY, eb.Rect.Width, eb.Rect.Height, rl.Black)
	rl.DrawText("X", eb.Rect.PosX, eb.Rect.PosY, eb.Rect.Height-2, rl.Black)
}

func (eb *ExitButton) NewPos(x, y int32) {
	eb.Rect.PosX = x
	eb.Rect.PosY = y
}

func (eb *ExitButton) IsDeleted() bool {
	return eb.isDeleted
}

func (eb *ExitButton) Delete() {
	eb.isDeleted = true
}
