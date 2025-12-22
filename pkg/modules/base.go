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
		return ei.GetLayer() < ej.GetLayer() // < вместо >
	})
	return es
}
