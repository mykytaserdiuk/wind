package modules

import (
	"fmt"

	"github.com/nikitaserdiuk9/pkg/models"
	"github.com/nikitaserdiuk9/pkg/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Panel struct {
	name string

	Width  int32
	Height int32
	PosX   int32
	PosY   int32

	BaseColor rl.Color

	Shrinking bool
	Dead      bool
	Color     rl.Color
}

func NewPanel(
	name string,
	width int32,
	height int32,
	posX int32,
	posY int32,
	color rl.Color,
) *Panel {

	return &Panel{
		name:      name,
		Width:     width,
		Height:    height,
		PosX:      posX,
		PosY:      posY,
		BaseColor: color,
		Color:     color,

		Shrinking: false,
		Dead:      false,
	}
}

func (b *Panel) Split(persent float32, horizontal bool) *Panel {
	var newBox *Panel
	if horizontal {
		diff := b.Height - int32(float32(b.Height)*persent)
		newBox = NewPanel(b.name+"_split", b.Width, diff, b.PosX, b.PosY+int32(float32(b.Height)*persent), b.BaseColor)
		b.Height = int32(float32(b.Height) * persent)
	} else {
		diff := b.Width - int32(float32(b.Width)*persent)
		newBox = NewPanel(b.name+"_split", diff, b.Height, b.PosX+int32(float32(b.Width)*persent), b.PosY, b.BaseColor)
		b.Width = int32(float32(b.Width) * persent)
	}

	return newBox
}

func (b *Panel) SetParent() {
	// Implementation for setting parent if needed
}
func (b *Panel) GetBounds() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(b.PosX),
		Y:      float32(b.PosY),
		Width:  float32(b.Width),
		Height: float32(b.Height),
	}
}
func (b *Panel) IsDead() bool {
	return b.Dead
}

func (b *Panel) Draw() {
	rl.DrawRectangle(b.PosX, b.PosY, b.Width, b.Height, b.Color)
	rl.DrawRectangleLines(b.PosX, b.PosY, b.Width, b.Height, rl.Black)
}

func (b *Panel) Update(dt float32) {
	if b.Shrinking {
		if b.Width > 0 && b.Height > 0 {
			b.Width -= int32(models.BoxDeleteSpeed * dt * 100)
			b.Height -= int32(models.BoxDeleteSpeed * dt * 100)
			// b.PosX += int32(models.BoxDeleteSpeed * dt * 100)
			// b.PosY -= int32(models.BoxDeleteSpeed * dt * 100)
		} else {
			b.Shrinking = false
			b.Dead = true
		}
	}

}

func (b *Panel) OnHover() {
	fmt.Println("box ", b.name, " hovered")
	if b.Shrinking || b.Dead {
		return
	}

	b.Color = utils.MakeLighter(b.BaseColor, 0.1)
}
func (b *Panel) OnUnhover() {
	fmt.Println("box ", b.name, " unhovered")
	if b.Shrinking || b.Dead {
		return
	}
	b.Color = b.BaseColor
}
func (b *Panel) OnClick(mouse rl.Vector2) {
	fmt.Println("box ", b.name, " clicked at ", mouse)
	if b.Shrinking || b.Dead {
		return
	}
	b.Color = rl.Red
	b.Shrinking = true
}
func (b *Panel) OnDrag(mouse rl.Vector2) {
	if b.Shrinking || b.Dead {
		return
	}

	fmt.Println("box ", b.name, " dragged at ", mouse)
}
func (b *Panel) OnDrop(mouse rl.Vector2) {
	if b.Shrinking || b.Dead {
		return
	}

	fmt.Println("box ", b.name, " dropped at ", mouse)
}
