package hud

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/pkg/models"
	"github.com/nikitaserdiuk9/pkg/modules"
)

type HUDPanel struct {
	BaseRect models.Rect
	Color    rl.Color

	elements []modules.HUDElement

	newElementCh chan modules.Element

	ColorViewer *ColorViewer
}

func NewHUDPanel(baseRect *models.Rect, color rl.Color, newElementCh chan modules.Element) *HUDPanel {
	return &HUDPanel{
		BaseRect:     *baseRect,
		Color:        color,
		newElementCh: newElementCh,
		elements:     []modules.HUDElement{NewColorViewer(0, 0)},
	}
}

func (hp *HUDPanel) GetBounds() rl.Rectangle {
	return hp.BaseRect.GetBounds()
}

func (hp *HUDPanel) Update(dt float32) {
	for _, el := range hp.elements {
		el.Update(dt)
	}
}

func (hp *HUDPanel) OnKeyInput(key int, pressed bool) {
	panic("not implemented") // TODO: Implement
}

func (hp *HUDPanel) OnLeftClick(mouse rl.Vector2) {
	fmt.Println("HUD left click ", mouse)
}

func (hp *HUDPanel) OnRightClick(mouse rl.Vector2) {
	fmt.Println("HUD right click ", mouse)
}
func (hp *HUDPanel) OnMouseWheel(value float32) {
	mousePos := rl.GetMousePosition()
	fmt.Println("1", mousePos)

	for _, el := range hp.elements {
		if rl.CheckCollisionPointRec(mousePos, el.GetBounds()) {
			el.OnMouseWheel(value)
			break
		}
	}
}
func (hp *HUDPanel) Draw() {
	rl.DrawRectangleRec(hp.BaseRect.GetBounds(), hp.Color)
	for _, el := range hp.elements {
		el.Draw()
	}
}
