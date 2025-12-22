package modules

import (
	"fmt"

	"github.com/nikitaserdiuk9/pkg/models"
	"github.com/nikitaserdiuk9/pkg/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Panel struct {
	name string

	BaseRect  *models.Rect
	BaseColor rl.Color
	baseLayer int8

	Shrinking bool
	Dead      bool

	dragging bool

	layer       int8
	Rect        models.Rect
	Color       rl.Color
	CloseButton *ExitButton
}

func NewPanel(
	name string,
	rect *models.Rect,
	color rl.Color,
) *Panel {

	baseRect := rect.Clone()
	closeButton := NewExitButton(models.NewRect(
		rect.PosX,
		rect.PosY,
		15,
		15,
	))
	return &Panel{
		layer:     1,
		baseLayer: 1,
		name:      name,
		BaseRect:  baseRect,
		Rect:      *rect,
		BaseColor: color,
		Color:     color,

		Shrinking:   false,
		Dead:        false,
		CloseButton: closeButton,
	}
}

func (p *Panel) Split(persent float32, horizontal bool) *Panel {
	var newPanel *Panel
	var rect *models.Rect
	if horizontal {
		diff := p.Rect.Height - int32(float32(p.Rect.Height)*persent)
		rect = models.NewRect(p.Rect.PosX, p.Rect.PosY+int32(float32(p.Rect.Height)*persent), p.Rect.Width, diff)
		p.Rect.Height = p.Rect.Height - diff
	} else {
		diff := p.Rect.Width - int32(float32(p.Rect.Width)*persent)
		rect = models.NewRect(p.Rect.PosX+int32(float32(p.Rect.Width)*persent), p.Rect.PosY, diff, p.Rect.Height)
		p.Rect.Width = p.Rect.Width - diff
	}

	newPanel = NewPanel(p.name+"_split", rect, p.BaseColor)
	p.CloseButton = NewExitButton(models.NewRect(
		p.Rect.PosX,
		p.Rect.PosY,
		15,
		15,
	))

	return newPanel
}

func (p *Panel) GetLayer() int8 {
	return p.layer
}

func (p *Panel) GetBounds() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(p.Rect.PosX),
		Y:      float32(p.Rect.PosY),
		Width:  float32(p.Rect.Width),
		Height: float32(p.Rect.Height),
	}
}
func (p *Panel) IsDead() bool {
	return p.Dead
}

func (p *Panel) Draw() {
	rl.DrawRectangle(p.Rect.PosX, p.Rect.PosY, p.Rect.Width, p.Rect.Height, p.Color)
	rl.DrawRectangleLines(p.Rect.PosX, p.Rect.PosY, p.Rect.Width, p.Rect.Height, rl.Black)

	if !p.CloseButton.IsDeleted() {
		p.CloseButton.Draw()
	}
}

func (p *Panel) Update(dt float32) {
	if p.dragging {
		mouse := rl.GetMousePosition()
		p.Rect.PosX = int32(mouse.X) - p.Rect.Width/2
		p.Rect.PosY = int32(mouse.Y) - p.Rect.Height/2

		p.CloseButton.NewPos(p.Rect.PosX, p.Rect.PosY)
		return
	}

	if p.Shrinking {
		if p.Rect.Width > 0 && p.Rect.Height > 0 {
			p.Rect.Width -= int32(models.BoxDeleteSpeed * dt * 100)
			p.Rect.Height -= int32(models.BoxDeleteSpeed * dt * 100)
		} else {
			p.Shrinking = false
			p.Dead = true
		}
		return
	}
}

func (p *Panel) OnHover() {
	fmt.Println("box ", p.name, " hovered")
	if p.Shrinking || p.Dead {
		return
	}

	p.Color = utils.MakeLighter(p.BaseColor, 0.1)
}
func (p *Panel) OnUnhover() {
	fmt.Println("box ", p.name, " unhovered")
	if p.Shrinking || p.Dead {
		return
	}
	p.Color = p.BaseColor
}

func (p *Panel) OnLeftClick(mouse rl.Vector2) {
	fmt.Println("box ", p.name, " clicked at ", mouse)
	if p.Shrinking || p.Dead {
		return
	}

	p.layer = 10

	if rl.CheckCollisionPointRec(mouse, p.CloseButton.GetBounds()) {
		fmt.Println("Closing box ", p.name)
		p.Shrinking = true
		p.CloseButton.Delete()
	}
}
func (p *Panel) OnRightClick(mouse rl.Vector2) {
	fmt.Println("box ", p.name, " RClicked at ", mouse)
	if p.Shrinking || p.Dead {
		return
	}
}

func (p *Panel) OnDrag(mouse rl.Vector2) {
	if p.Shrinking || p.Dead {
		return
	}

	p.dragging = true
	fmt.Println("drag box ", p.name, " layer ", p.layer)
}
func (p *Panel) OnDrop(mouse rl.Vector2) {
	if p.Shrinking || p.Dead {
		return
	}

	p.dragging = false
	p.layer = p.baseLayer
	fmt.Println("drop box ", p.name, " layer ", p.layer)
}
