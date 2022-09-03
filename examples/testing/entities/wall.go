package entities

import (
	"image/color"

	"github.com/odedro987/gixel-engine/gixel"
	"github.com/odedro987/gixel-engine/gixel/graphic"
	"github.com/odedro987/gixel-engine/gixel/systems/physics"
)

type Wall struct {
	gixel.BaseGxlSprite
	// Systems
	physics.Physics
}

func NewWall(x, y float64) *Wall {
	w := &Wall{}
	w.SetPosition(x, y)

	return w
}

func (w *Wall) Init(game *gixel.GxlGame) {
	w.BaseGxlSprite.Init(game)
	w.Physics.Init(w)

	w.ApplyGraphic(game.Graphics().MakeGraphic(32, 32, color.White, graphic.CacheOptions{}))

	w.Color().R = 0
	w.Color().G = 0
}

func (w *Wall) Update(elapsed float64) error {
	err := w.BaseGxlSprite.Update(elapsed)
	if err != nil {
		return err
	}

	w.Physics.Update(elapsed)

	return nil
}
