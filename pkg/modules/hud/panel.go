package hud

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/pkg/models"
	"github.com/nikitaserdiuk9/pkg/modules"
)

type HUDPanel struct {
	BaseRect models.Rect
	Color    rl.Color

	elements []modules.HUDElement

	hovered      modules.HUDElement
	newElementCh chan modules.Element
}

func NewHUDPanel(baseRect *models.Rect, color rl.Color, newElementCh chan modules.Element) *HUDPanel {
	colorViewer := NewColorViewer(0, 0)
	panCreatorRect := models.NewRect(baseRect.Width+baseRect.PosX-100, baseRect.PosY, 100, baseRect.Height)
	hud := &HUDPanel{
		BaseRect:     *baseRect,
		Color:        color,
		newElementCh: newElementCh,
	}

	panCreatorClickHandler := func() {
		color := colorViewer.GetColor()
		newPanelRect := models.NewRect(0, 0, 250, 250)
		newPanel := modules.NewPanel(time.Now().GoString(), newPanelRect, color, hud.newElementCh)
		hud.newElementCh <- newPanel
	}

	hud.AddElement(colorViewer)
	hud.AddElement(NewPanelCreator(rl.Blue, *panCreatorRect, panCreatorClickHandler))
	return hud
}

func (hp *HUDPanel) OnHover(mousePos rl.Vector2) {
	for _, el := range hp.elements {
		if rl.CheckCollisionPointRec(mousePos, el.GetBounds()) {
			hp.hovered = el
			el.OnHover(mousePos)
		}
	}
}

func (hp *HUDPanel) OnUnhover(mousePos rl.Vector2) {
	for _, el := range hp.elements {
		if rl.CheckCollisionPointRec(mousePos, el.GetBounds()) || hp.hovered == el {
			el.OnUnhover(mousePos)
		}
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
	for _, el := range hp.elements {
		if rl.CheckCollisionPointRec(mouse, el.GetBounds()) {
			el.OnLeftClick(mouse)
		}
	}
}

func (hp *HUDPanel) OnRightClick(mouse rl.Vector2) {
	for _, el := range hp.elements {
		if rl.CheckCollisionPointRec(mouse, el.GetBounds()) {
			el.OnRightClick(mouse)
		}
	}
}
func (hp *HUDPanel) OnMouseWheel(value float32) {
	mousePos := rl.GetMousePosition()

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

func (hp *HUDPanel) AddElement(el modules.HUDElement) {
	hp.elements = append(hp.elements, el)
}
