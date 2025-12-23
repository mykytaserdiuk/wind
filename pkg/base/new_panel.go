package base

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/pkg/models"
)

type NewPanel struct {
	BaseRect models.Rect
	Color    rl.Color

	ColorViewer *ColorViewer
}

func NewNewPanel(baseRect *models.Rect, color rl.Color) *NewPanel {
	return &NewPanel{
		BaseRect: *baseRect,
		Color:    color,
	}
}

func (np *NewPanel) Draw() {
	// each := np.BaseRect.Width / 3
	// rl.DrawRectangleLines(np.BaseRect.PosX, np.BaseRect.PosY, each, np.BaseRect.Height, rl.Black)
	// rl.DrawRectangleLines(np.BaseRect.PosX+each, np.BaseRect.PosY, each, np.BaseRect.Height, rl.Black)
	// rl.DrawRectangleLines(np.BaseRect.PosX+2*each, np.BaseRect.PosY, each, np.BaseRect.Height, rl.Black)

	rl.DrawRectangleRec(np.BaseRect.GetBounds(), np.Color)
}
