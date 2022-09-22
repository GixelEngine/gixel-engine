//go:build !debug

package debug

import (
	"github.com/GixelEngine/gixel-engine/gixel/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Collision struct{}

func (c *Collision) Update() {}

func (c *Collision) DrawBounds(screen *ebiten.Image, bounds math.GxlRectangle) {}
