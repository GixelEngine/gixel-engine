package main

import (
	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type movementRequirements interface {
	Velocity() *math.GxlPoint
	MaxVelocity() *math.GxlPoint
	Acceleration() *math.GxlPoint
	Drag() *math.GxlPoint
	Facing() *gixel.GxlDirection
}

type Movement struct {
	speed     float64
	maxSpeed  float64
	canMove   bool
	isMoving  bool
	moveAngle float64
	animName  string
	upKey     ebiten.Key
	downKey   ebiten.Key
	rightKey  ebiten.Key
	leftKey   ebiten.Key
	subject   *movementRequirements
}

func (p *Movement) Init(subject movementRequirements) {
	p.subject = &subject
	p.canMove = true
	p.isMoving = false
	p.speed = 0
	p.maxSpeed = 140
	(*p.subject).Drag().Set(1000, 1000)
	(*p.subject).Acceleration().X = 0
}

func (p *Movement) Update(elapsed float64) {
	p.keyboardControl()
	p.updateFacing()
	p.updateMovement(elapsed)
}

func (p *Movement) SetControls(up, down, left, right ebiten.Key) {
	p.upKey, p.downKey, p.leftKey, p.rightKey = up, down, left, right
}

func (p *Movement) CanMove() *bool {
	return &p.canMove
}

func (p *Movement) GetAnimName() string {
	return p.animName
}

func (p *Movement) keyboardControl() {
	up := ebiten.IsKeyPressed(p.upKey)
	down := ebiten.IsKeyPressed(p.downKey)
	left := ebiten.IsKeyPressed(p.leftKey)
	right := ebiten.IsKeyPressed(p.rightKey)

	if !p.canMove {
		p.isMoving = false
		return
	}

	if up || down || left || right {
		if up && down {
			up = false
			down = false
		}
		if left && right {
			left = false
			right = false
		}
		if !up && !down && !left && !right {
			p.isMoving = false
			return
		}

		newAngle := 0.0

		if up {
			newAngle = -90
			if left {
				newAngle -= 45
			} else if right {
				newAngle += 45
			}
		} else if down {
			newAngle = 90
			if left {
				newAngle += 45
			} else if right {
				newAngle -= 45
			}
		} else if left {
			newAngle = 180
		} else if right {
			newAngle = 0
		}

		p.move(p.maxSpeed, newAngle)

		return
	}

	p.isMoving = false
}

func (p *Movement) move(newSpeed float64, newAngle float64) {
	if !p.canMove {
		return
	}

	p.moveAngle = newAngle
	p.speed = newSpeed
	p.isMoving = true
}

func (p *Movement) updateMovement(elapsed float64) {
	isFront := p.moveAngle >= -25 && p.moveAngle <= 180 || p.moveAngle >= -180 && p.moveAngle <= -155

	if p.isMoving {
		(*p.subject).Acceleration().X += p.speed*0.15 + p.speed*elapsed
		if (*p.subject).Acceleration().X >= p.speed {
			(*p.subject).Acceleration().X = p.speed
		}

		if (p.moveAngle > -45 && p.moveAngle < 45) || (p.moveAngle > 135 && p.moveAngle <= 180) || (p.moveAngle < -135 && p.moveAngle >= -180) {
			p.animName = "WalkSide"
		} else {
			if isFront {
				p.animName = "WalkFront"
			} else {
				p.animName = "WalkBack"
			}
		}

		(*p.subject).Velocity().Set((*p.subject).Acceleration().X, 0)
		newVel := (*p.subject).Velocity().PivotDegrees(&math.GxlPoint{X: 0, Y: 0}, p.moveAngle)
		(*p.subject).Velocity().Set(newVel.X, newVel.Y)

		return
	}

	(*p.subject).Acceleration().X = 0

	if isFront {
		p.animName = "StandFront"
	} else {
		p.animName = "StandBack"
	}
}

func (p *Movement) updateFacing() {
	*(*p.subject).Facing() = 0

	if p.moveAngle >= -25 && p.moveAngle <= 25 {
		*(*p.subject).Facing() |= gixel.Right
	} else if (p.moveAngle >= -180 && p.moveAngle <= -155) || (p.moveAngle >= 155 && p.moveAngle <= 180) {
		*(*p.subject).Facing() |= gixel.Left
	}
}
