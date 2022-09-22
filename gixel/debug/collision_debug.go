//go:build debug

package debug

import (
	"image/color"

	"github.com/GixelEngine/gixel-engine/gixel/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Collision struct {
	enabled bool
}

func (c *Collision) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF10) {
		c.enabled = !c.enabled
	}
}

func (c *Collision) DrawBounds(screen *ebiten.Image, bounds math.GxlRectangle) {
	if !c.enabled {
		return
	}
	ebitenutil.DrawRect(screen, *bounds.X(), *bounds.Y(), *bounds.W(), *bounds.H(), color.RGBA{R: 127, G: 0, B: 0, A: 127})
}
