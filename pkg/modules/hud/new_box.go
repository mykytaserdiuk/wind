package hud

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/pkg/models"
	"github.com/nikitaserdiuk9/pkg/utils"
)

type PanelCreator struct {
	rect models.Rect

	isHover bool
	onClick func()

	baseColor rl.Color
	color     rl.Color
}

func NewPanelCreator(color rl.Color, rect models.Rect, onClick func()) *PanelCreator {
	return &PanelCreator{
		baseColor: color,
		onClick:   onClick,
		color:     color,
		rect:      rect,
		isHover:   false,
	}
}

func (c *PanelCreator) Draw() {
	rl.DrawRectangle(c.rect.PosX, c.rect.PosY, c.rect.Width, c.rect.Height, c.color)
	rl.DrawRectangleLines(c.rect.PosX, c.rect.PosY, c.rect.Width, c.rect.Height, rl.Black)
}

func (c *PanelCreator) Update(dt float32) {
	if c.isHover {
		c.color = utils.MakeLighter(c.baseColor, 0.1)
	} else {
		c.color = c.baseColor
	}
}
func (c *PanelCreator) OnKeyInput(key int, pressed bool) {
	panic("not implemented") // TODO: Implement
}

func (c *PanelCreator) OnHover(mousePos rl.Vector2) {
	fmt.Println("Panel creator Hovered")
	c.isHover = true
}

func (c *PanelCreator) OnUnhover(mousePos rl.Vector2) {
	fmt.Println("Panel creator UnHovered")

	c.isHover = false
}

func (c *PanelCreator) GetLayer() int8 {
	panic("not implemented") // TODO: Implement
}

func (c *PanelCreator) GetBounds() rl.Rectangle {
	return c.rect.GetBounds()
}

func (c *PanelCreator) OnMouseWheel(value float32) {
	return
}

func (c *PanelCreator) OnLeftClick(mouse rl.Vector2) {
	c.onClick()
}

func (c *PanelCreator) OnRightClick(mouse rl.Vector2) {
	panic("not implemented") // TODO: Implement
}
