package main

import (
	"github.com/nikitaserdiuk9/pkg/base"
	"github.com/nikitaserdiuk9/pkg/models"
	"github.com/nikitaserdiuk9/pkg/modules"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	app = &base.Application{}
)

func init() {
	app.Init()
}

func main() {
	rl.InitWindow(models.WindowWidth, models.WindowHeight, "WIND")
	defer rl.CloseWindow()
	box1 := modules.NewPanel("box1", 200, 200, 10, 10, rl.Blue)
	box2 := box1.Split(0.5, true)
	app.AddElement(box1)
	app.AddElement(box2)

	for app.IsActive() {
		app.Input()

		app.Update()

		app.Render()
	}

	app.Close()
}
