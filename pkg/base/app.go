package base

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/pkg/models"
	"github.com/nikitaserdiuk9/pkg/modules"
	"github.com/nikitaserdiuk9/pkg/utils"
)

type Application struct {
	active bool
	cam    rl.Camera2D

	elements     modules.Elements
	newElementCh chan modules.Element
	newPanel     *NewPanel

	hovered       modules.Element
	activeElement modules.Element
	dragging      bool
	middleBtnDown bool
}

func NewApplication() *Application {
	return &Application{
		active:   false,
		newPanel: &NewPanel{},
		cam:      rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(0, 0), 0, 1),
		elements: modules.Elements{},
		hovered:  nil,
		dragging: false,
	}
}

func (app *Application) Init() {
	app.elements = []modules.Element{}
	app.hovered = nil
	app.activeElement = nil
	app.dragging = false

	newPanelRect := models.NewRect(app.newPanel.BaseRect.PosX, app.newPanel.BaseRect.PosY, app.newPanel.BaseRect.Width, 100)
	app.newPanel = NewNewPanel(newPanelRect, rl.LightGray)

	rl.SetTargetFPS(60)
}

func (app *Application) GetNewElementChannel() chan modules.Element {
	newElementCh := make(chan modules.Element)
	app.newElementCh = newElementCh

	go func() {
		for {
			newElem, ok := <-app.newElementCh
			if !ok {
				fmt.Println("New element channel closed")
				return
			}

			app.AddElement(newElem)
		}
	}()

	return app.newElementCh
}

func (app *Application) IsActive() bool {
	app.active = rl.WindowShouldClose() == false
	return app.active
}

func (app *Application) Update() {
	dt := rl.GetFrameTime()

	if app.middleBtnDown {
		mouseDt := rl.GetMouseDelta()
		app.cam.Target = rl.NewVector2(app.cam.Target.X-mouseDt.X, app.cam.Target.Y-mouseDt.Y)
	}

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
	mouseScreen := rl.GetMousePosition()
	mouseWorld := rl.GetScreenToWorld2D(mouseScreen, app.cam)

	newHovered := app.findHovered(mouseWorld)
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
		app.activeElement.OnLeftClick(mouseWorld)
		app.elements = app.elements.LayerSort()
	}
	if rl.IsMouseButtonPressed(rl.MouseRightButton) && app.hovered != nil {
		app.activeElement = app.hovered
		app.activeElement.OnRightClick(mouseWorld)
		app.elements = app.elements.LayerSort()
	}
	if rl.IsMouseButtonPressed(rl.MouseMiddleButton) {
		app.middleBtnDown = true
	}
	if rl.IsMouseButtonReleased(rl.MouseMiddleButton) {
		app.middleBtnDown = false
	}
	if rl.IsKeyDown(rl.KeyR) {
		app.cam.Zoom = 1
		app.cam.Target = rl.NewVector2(0, 0)
	}

	mouseMove := rl.GetMouseWheelMove()
	if mouseMove != 0 {
		mltp := float32(0.03)
		if rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift) ||
			rl.IsKeyDown(rl.KeyLeftAlt) || rl.IsKeyDown(rl.KeyRightAlt) {
			mltp = 0.1
		}
		app.cam.Zoom = utils.Clamp(app.cam.Zoom+(mouseMove*mltp), 0.5, 1.2)
	}

	// Drag
	if app.dragging && app.activeElement != nil {
		app.activeElement.OnDrag(mouseWorld)
	}

	// Mouse release
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		if app.dragging && app.activeElement != nil {
			app.activeElement.OnDrop(mouseWorld)
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
	rl.BeginMode2D(app.cam)

	rl.ClearBackground(rl.RayWhite)

	for _, element := range app.elements {
		element.Draw()
	}
	rl.DrawText(fmt.Sprintf("%.2f", rl.GetFrameTime()*100), 500, 500, 20, rl.Green)
	app.newPanel.Draw()

	rl.EndMode2D()
	rl.EndDrawing()
}

func (app *Application) findHovered(mouse rl.Vector2) modules.Element {
	mouseScreen := rl.GetMousePosition()
	mouseWorld := rl.GetScreenToWorld2D(mouseScreen, app.cam)

	for i := len(app.elements) - 1; i >= 0; i-- {
		if rl.CheckCollisionPointRec(mouseWorld, app.elements[i].GetBounds()) {
			return app.elements[i]
		}
	}

	return nil
}
