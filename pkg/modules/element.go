package modules

import (
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Element interface {
	Draw()
	Update(dt float32)
	GetLayer() int8

	GetBounds() rl.Rectangle
	// Initialize()
	IsDead() bool

	OnHover()
	OnUnhover()

	OnLeftClick(mouse rl.Vector2)
	OnRightClick(mouse rl.Vector2)
	OnDrag(mouse rl.Vector2)
	OnDrop(mouse rl.Vector2)
}

type Elements []Element

func (es Elements) LayerSort() Elements {
	sort.Slice(es, func(i, j int) bool {
		ei, ej := es[i], es[j]
		return ei.GetLayer() < ej.GetLayer()
	})
	return es
}

type HUD interface {
	Draw()
	Update(dt float32)
	OnKeyInput(key int, pressed bool)

	GetBounds() rl.Rectangle
	OnMouseWheel(value float32)
	AddElement(el HUDElement)

	OnHover(mousePos rl.Vector2)
	OnUnhover(mousePos rl.Vector2)

	OnLeftClick(mouse rl.Vector2)
	OnRightClick(mouse rl.Vector2)
}

type HUDElement interface {
	Draw()
	Update(dt float32)
	GetLayer() int8

	GetBounds() rl.Rectangle
	OnMouseWheel(value float32)

	OnHover(mousePos rl.Vector2)
	OnUnhover(mousePos rl.Vector2)

	OnLeftClick(mouse rl.Vector2)
	OnRightClick(mouse rl.Vector2)
}
