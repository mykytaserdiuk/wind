package base

import (
	"fmt"

	"github.com/nikitaserdiuk9/pkg/modules"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Application struct {
	active bool

	elements []modules.Element

	hovered       int
	activeElement int
	dragging      bool
}

func (app *Application) IsActive() bool {
	app.active = rl.WindowShouldClose() == false
	return app.active
}

func (app *Application) Init() {
	app.elements = []modules.Element{}
	app.hovered = -1
	app.activeElement = -1
	app.dragging = false

	rl.SetTargetFPS(60)
}

func (app *Application) Update() {
	dt := rl.GetFrameTime()
	for i := 0; i < len(app.elements); {
		el := app.elements[i]
		el.Update(dt)

		if el.IsDead() {
			fmt.Println("Removing element at index ", i)
			app.elements = append(app.elements[:i], app.elements[i+1:]...)
		} else {
			i++
		}
	}
}

func (app *Application) Input() {
	mouse := rl.GetMousePosition()

	newHovered := app.findHovered(mouse)

	// Hover / Unhover
	if newHovered != app.hovered {
		if app.hovered != -1 {
			app.elements[app.hovered].OnUnhover()
		}
		if newHovered != -1 {
			app.elements[newHovered].OnHover()
		}
		app.hovered = newHovered
	}

	// Mouse press
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && app.hovered != -1 {
		app.activeElement = app.hovered
		app.dragging = true
		app.elements[app.activeElement].OnClick(mouse)
	}

	// Drag
	if app.dragging && app.activeElement != -1 {
		app.elements[app.activeElement].OnDrag(mouse)
	}

	// Mouse release
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		if app.dragging && app.activeElement != -1 {
			app.elements[app.activeElement].OnDrop(mouse)
		}
		app.dragging = false
		app.activeElement = -1
	}
}
func (app *Application) Close() {
	app.elements = nil
}

func (app *Application) AddElement(element modules.Element) {
	app.elements = append(app.elements, element)
}

func (app *Application) Render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.DrawText(fmt.Sprintf("%.2f", rl.GetFrameTime()*100), 500, 500, 20, rl.Green)

	for _, element := range app.elements {
		element.Draw()
	}

	rl.EndDrawing()

}

func (app *Application) findHovered(mouse rl.Vector2) int {
	for i, el := range app.elements {
		if rl.CheckCollisionPointRec(mouse, el.GetBounds()) {
			return i
		}
	}
	return -1
}
