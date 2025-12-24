package hud

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/pkg/models"
)

type ColorViewer struct {
	rect    models.Rect
	cells   [3]rl.Rectangle // R, G, B
	preview rl.Rectangle

	input [3]uint8 // R, G, B
	layer int8
}

func NewColorViewer(x, y int32) *ColorViewer {
	cv := &ColorViewer{
		// rect:   rl.Rectangle{X: x, Y: y, Width: 180, Height: 40},
		rect:  *models.NewRect(x, y, 180, 40),
		layer: 100,
	}
	for i := 0; i < 3; i++ {
		cellW := float32(50)
		r := rl.Rectangle{
			X:      float32(cv.rect.PosX) + float32(i)*(cellW+5),
			Y:      float32(cv.rect.PosY),
			Width:  cellW,
			Height: float32(cv.rect.Height),
		}
		cv.cells[i] = r
	}
	return cv
}

func (c *ColorViewer) Draw() {
	labels := []string{"R", "G", "B"}
	for i := 0; i < 3; i++ {
		r := c.cells[i]

		bg := rl.LightGray

		rl.DrawRectangleRec(r, bg)
		rl.DrawRectangleLinesEx(r, 1, rl.DarkGray)

		rl.DrawText(labels[i], int32(r.X+4), int32(r.Y-14), 10, rl.DarkGray)
		rl.DrawText(fmt.Sprintf("%d", c.input[i]), int32(r.X), int32(r.Y), 14, rl.Black)

	}

	// preview
	rl.DrawRectangle(
		int32(c.rect.PosX+170),
		int32(c.rect.PosY),
		30,
		30,
		rl.Color{R: c.input[0], G: c.input[1], B: c.input[2], A: 255},
	)
}

func (c *ColorViewer) GetBounds() rl.Rectangle {
	return c.rect.GetBounds()
}

func (c *ColorViewer) GetLayer() int8 {
	return c.layer
}

func (c *ColorViewer) OnLeftClick(mouse rl.Vector2) {}
func (c *ColorViewer) OnRightClick(mouse rl.Vector2) {

}
func (c *ColorViewer) OnMouseWheel(value float32) {
	mousePos := rl.GetMousePosition()
	for i, cell := range c.cells {
		if rl.CheckCollisionPointRec(mousePos, cell) {
			c.input[i] += uint8(value)

			fmt.Println(c.input)
			break
		}
	}
}
func (c *ColorViewer) Update(dt float32) {}
