package modules

import rl "github.com/gen2brain/raylib-go/raylib"

type Element interface {
	Draw()
	Update(dt float32)
	SetParent()

	GetBounds() rl.Rectangle
	// Initialize()
	IsDead() bool

	OnHover()
	OnUnhover()

	OnClick(mouse rl.Vector2)
	OnDrag(mouse rl.Vector2)
	OnDrop(mouse rl.Vector2)
}
