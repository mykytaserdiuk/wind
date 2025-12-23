package main

import (
	"github.com/nikitaserdiuk9/pkg/base"
	"github.com/nikitaserdiuk9/pkg/models"
	"github.com/nikitaserdiuk9/pkg/modules"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	app = base.NewApplication()
)

func init() {
	app.Init()
}

func main() {
	rl.InitWindow(models.WindowWidth, models.WindowHeight, "WIND")
	defer rl.CloseWindow()

	rect := models.NewRect(10, 10, 500, 200)
	newElemCh := app.GetNewElementChannel()

	box1 := modules.NewPanel("box1", rect, rl.Blue, newElemCh)
	box2 := box1.Split(0.5, false)

	app.AddElement(box1)
	app.AddElement(box2)

	for app.IsActive() {
		app.Input()

		app.Update()

		app.Render()
	}

	app.Close()
}
