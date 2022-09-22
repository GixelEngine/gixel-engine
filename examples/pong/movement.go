package main

import (
	"github.com/GixelEngine/gixel-engine/gixel/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type imports interface {
	Velocity() *math.GxlPoint
	Y() *float64
	H() *int
}

type Movement struct {
	upKey, downKey ebiten.Key
	speed          float64
	subject        *imports
}

func (p *Movement) Init(subject imports) {
	p.subject = &subject
	p.speed = 800
}

func (p *Movement) Update(elapsed float64) {
	p.keyboardControl()
}

func (p *Movement) SetControls(upKey, downKey ebiten.Key) {
	p.upKey, p.downKey = upKey, downKey
}

func (p *Movement) keyboardControl() {
	up := ebiten.IsKeyPressed(p.upKey)
	down := ebiten.IsKeyPressed(p.downKey)

	vel := 0.0

	subjY := *(*p.subject).Y()
	subjH := float64(*(*p.subject).H())

	if up && subjY > 0 {
		vel += -p.speed
	}

	if down && subjY < GAME_HEIGHT-subjH {
		vel += p.speed
	}

	(*p.subject).Velocity().Y = vel
}
