package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/odedro987/gixel-engine/gixel"
	"github.com/odedro987/gixel-engine/gixel/graphic"
	"github.com/odedro987/gixel-engine/gixel/systems/physics"
)

const BALL_SPEED = 800
const MAX_ANGLE = 45

type Ball struct {
	gixel.BaseGxlSprite

	active bool

	// Systems
	physics.Physics
}

func NewBall() *Ball {
	b := &Ball{
		active: false,
	}

	return b
}

func (b *Ball) Init(game *gixel.GxlGame) {
	b.BaseGxlSprite.Init(game)
	b.Physics.Init(b)

	b.ApplyGraphic(game.Graphics().MakeGraphic(32, 32, color.White, graphic.CacheOptions{}))
	*b.Visible() = false
}

func (b *Ball) Spawn() {
	b.SetPosition(GAME_WIDTH/2-16, GAME_HEIGHT/2-16)
	*b.Visible() = true
	b.active = true

	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	b.Velocity().X = float64(BALL_SPEED)

	if randGen.Float64() > 0.5 {
		b.Velocity().X *= -1
	}

	b.Velocity().Y = float64(BALL_SPEED) * (randGen.Float64() - 0.5)
}

func (b *Ball) Active() *bool {
	return &b.active
}

func (b *Ball) FlipVertical() {
	b.Velocity().Y *= -1
}

func (b *Ball) FlipHorizontal(normal float64, dir float64) {
	b.Velocity().X = BALL_SPEED * math.Cos((MAX_ANGLE*(math.Pi/180))*normal) * dir

	if dir == 1 {
		b.Velocity().Y = BALL_SPEED * -math.Sin((MAX_ANGLE*(math.Pi/180))*normal)
	} else {
		b.Velocity().Y = BALL_SPEED * math.Sin((MAX_ANGLE*(math.Pi/180))*normal) * dir
	}
}

func (b *Ball) Update(elapsed float64) error {
	if !b.active {
		return nil
	}

	err := b.BaseGxlSprite.Update(elapsed)
	if err != nil {
		return err
	}

	b.Physics.Update(elapsed)

	if *b.Y() > GAME_HEIGHT-float64(*b.H()) || *b.Y() < 0 {
		b.FlipVertical()
	}

	return nil
}
