package modules

import rl "github.com/gen2brain/raylib-go/raylib"

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
	// OnRightClick(mouse rl.Vector2)
	OnDrag(mouse rl.Vector2)
	OnDrop(mouse rl.Vector2)
}

type Elements []Element

func (es Elements) LayerSort() Elements {
	sorted := make(Elements, len(es))

	for i, element := range es {
		layer := element.GetLayer()
		if layer < 0 {
			sorted[i] = element
		} else {
			for j := i - 1; j >= 0 && sorted[j].GetLayer() > layer; j-- {
				sorted[j+1] = sorted[j]
			}
			sorted[i] = element
		}
	}

	return sorted
}
