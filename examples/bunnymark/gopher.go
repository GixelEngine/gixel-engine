package main

import (
	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/graphic"
	"github.com/GixelEngine/gixel-engine/gixel/systems/physics"
)

const GRAVITY = 10

type Gopher struct {
	gixel.BaseGxlSprite
	physics.Physics
}

func NewGopher(x, y float64) *Gopher {
	g := &Gopher{}
	g.SetPosition(x, y)
	return g
}

func (g *Gopher) Init(game *gixel.GxlGame) {
	g.BaseGxlSprite.Init(game)
	g.Physics.Init(g)

	g.Velocity().X = RangeFloat(0, 300)
	g.Velocity().Y = RangeFloat(1, 50)

	g.ApplyGraphic(game.Graphics().LoadGraphic("assets/gopher.png", graphic.CacheOptions{}))
}

func (g *Gopher) Update(elapsed float64) error {
	err := g.BaseGxlSprite.Update(elapsed)
	if err != nil {
		return err
	}

	g.Velocity().Y += GRAVITY

	if *g.X()+float64(g.Graphic().FrameWidth()) > float64(g.Game().W()) {
		g.Velocity().X *= -1
		*g.X() = float64(g.Game().W()) - float64(g.Graphic().FrameWidth())
	}
	if *g.X() < 0 {
		g.Velocity().X *= -1
		*g.X() = 0
	}
	if *g.Y()+float64(g.Graphic().FrameHeight()) > float64(g.Game().H()) {
		g.Velocity().Y *= -0.95
		*g.Y() = float64(g.Game().H()) - float64(g.Graphic().FrameHeight())
		if Chance(0.5) {
			g.Velocity().Y -= RangeFloat(0, 0.009)
		}
	}
	if *g.Y() < 0 {
		g.Velocity().Y = 0
		*g.Y() = 0
	}

	g.Physics.Update(elapsed)

	return nil
}
