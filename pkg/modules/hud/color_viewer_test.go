package hud

import (
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/pkg/models"
	"github.com/nikitaserdiuk9/pkg/modules"
)

func TestColorViewer_InitializationAndGetters(t *testing.T) {
	cv := NewColorViewer(5, 7)

	if cv.GetBounds().X != 5 || cv.GetBounds().Y != 7 {
		t.Fatalf("unexpected GetBounds() pos: got (%b,%b)", cv.GetBounds().X, cv.GetBounds().Y)
	}
	if cv.GetBounds().Width != 180 || cv.GetBounds().Height != 40 {
		t.Fatalf("unexpected GetBounds() size: got (%b,%b)", cv.GetBounds().Width, cv.GetBounds().Height)
	}

	expectedInput := [3]uint8{255, 255, 255}
	if cv.GetColor() != rl.NewColor(expectedInput[0], expectedInput[1], expectedInput[2], 255) {
		t.Fatalf("expected input %v, got %v", expectedInput, cv.GetColor())
	}

	if cv.GetLayer() != 100 {
		t.Fatalf("expected layer 100, got %d", cv.GetLayer())
	}

	col := cv.GetColor()
	if col.R != 255 || col.G != 255 || col.B != 255 || col.A != 255 {
		t.Fatalf("unexpected color: %#v", col)
	}
}

func TestHUDPanel_NewAndAddElement(t *testing.T) {
	base := models.NewRect(0, 0, 200, 100)
	ch := make(chan modules.Element, 4)

	hp := NewHUDPanel(base, rl.White, ch)
	if hp == nil {
		t.Fatalf("NewHUDPanel returned nil")
	}

	// NewHUDPanel should add the color viewer and the panel creator (2 elements)
	if len(hp.elements) != 2 {
		t.Fatalf("expected 2 elements after NewHUDPanel, got %d", len(hp.elements))
	}

	// Add a dummy element and ensure it appends
	d := &dummyHUDElement{layer: 1}
	hp.AddElement(d)
	if len(hp.elements) != 3 {
		t.Fatalf("expected 3 elements after AddElement, got %d", len(hp.elements))
	}
}

// dummyHUDElement satisfies modules.HUDElement for test purposes.
type dummyHUDElement struct {
	layer int8
}

func (d *dummyHUDElement) Draw()             {}
func (d *dummyHUDElement) Update(dt float32) {}
func (d *dummyHUDElement) GetLayer() int8    { return d.layer }
func (d *dummyHUDElement) GetBounds() rl.Rectangle {
	return rl.Rectangle{X: 0, Y: 0, Width: 10, Height: 10}
}
func (d *dummyHUDElement) OnMouseWheel(value float32)    {}
func (d *dummyHUDElement) OnHover(mousePos rl.Vector2)   {}
func (d *dummyHUDElement) OnUnhover(mousePos rl.Vector2) {}
func (d *dummyHUDElement) OnLeftClick(mouse rl.Vector2)  {}
func (d *dummyHUDElement) OnRightClick(mouse rl.Vector2) {}
