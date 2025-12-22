package base

import (
	"fmt"

	"github.com/nikitaserdiuk9/pkg/modules"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Application struct {
	active bool

	elements modules.Elements

	hovered       modules.Element
	activeElement modules.Element
	dragging      bool
}

func (app *Application) IsActive() bool {
	app.active = rl.WindowShouldClose() == false
	return app.active
}

func (app *Application) Init() {
	app.elements = []modules.Element{}
	app.hovered = nil
	app.activeElement = nil
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
		if app.hovered != nil {
			app.hovered.OnUnhover()
		}
		if newHovered != nil {
			newHovered.OnHover()
		}
		app.hovered = newHovered
	}

	// Mouse press
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && app.hovered != nil {
		app.activeElement = app.hovered
		app.dragging = true
		app.activeElement.OnLeftClick(mouse)
		app.elements = app.elements.LayerSort()
	}
	if rl.IsMouseButtonPressed(rl.MouseRightButton) && app.hovered != nil {
		app.activeElement = app.hovered
		app.activeElement.OnRightClick(mouse)
		app.elements = app.elements.LayerSort()
	}

	// Drag
	if app.dragging && app.activeElement != nil {
		app.activeElement.OnDrag(mouse)
	}

	// Mouse release
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		if app.dragging && app.activeElement != nil {
			app.activeElement.OnDrop(mouse)
			app.elements = app.elements.LayerSort()
		}
		app.dragging = false
		app.activeElement = nil
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

	for _, element := range app.elements {
		element.Draw()
	}
	rl.DrawText(fmt.Sprintf("%.2f", rl.GetFrameTime()*100), 500, 500, 20, rl.Green)

	rl.EndDrawing()
}

func (app *Application) findHovered(mouse rl.Vector2) modules.Element {
	for i := len(app.elements) - 1; i >= 0; i-- {
		if rl.CheckCollisionPointRec(mouse, app.elements[i].GetBounds()) {
			return app.elements[i]
		}
	}

	return nil
}
