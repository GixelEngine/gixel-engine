package physics

import (
	"github.com/GixelEngine/gixel-engine/gixel/math"
)

type imports interface {
	X() *float64
	Y() *float64
	Angle() *float64
}

type Exports interface {
	Velocity() *math.GxlPoint
	AngularVelocity() *float64
	AngularAcceleration() *float64
	AngularDrag() *float64
	MaxAngular() *float64
	MaxVelocity() *math.GxlPoint
	Acceleration() *math.GxlPoint
	Drag() *math.GxlPoint
}

type Physics struct {
	velocity            math.GxlPoint
	angularVelocity     float64
	angularAcceleration float64
	angularDrag         float64
	maxAngular          float64
	maxVelocity         math.GxlPoint
	drag                math.GxlPoint
	acceleration        math.GxlPoint
	subject             *imports
}

func (p *Physics) Init(subject imports) {
	p.subject = &subject
}

func (p *Physics) Update(elapsed float64) {
	velocityDelta := 0.5 * (computeVelocity(p.angularVelocity, p.angularAcceleration, p.angularDrag, p.maxAngular, elapsed) - p.angularVelocity)
	p.angularVelocity += velocityDelta
	*(*p.subject).Angle() += p.angularVelocity * elapsed
	p.angularVelocity += velocityDelta

	velocityDelta = 0.5 * (computeVelocity(p.velocity.X, p.acceleration.X, p.drag.X, p.maxVelocity.X, elapsed) - p.velocity.X)
	p.velocity.X += velocityDelta
	delta := p.velocity.X * elapsed
	p.velocity.X += velocityDelta
	*(*p.subject).X() += delta

	velocityDelta = 0.5 * (computeVelocity(p.velocity.Y, p.acceleration.Y, p.drag.Y, p.maxVelocity.Y, elapsed) - p.velocity.Y)
	p.velocity.Y += velocityDelta
	delta = p.velocity.Y * elapsed
	p.velocity.Y += velocityDelta
	*(*p.subject).Y() += delta
}

func (p *Physics) Velocity() *math.GxlPoint {
	return &p.velocity
}

func (p *Physics) AngularVelocity() *float64 {
	return &p.angularVelocity
}

func (p *Physics) AngularAcceleration() *float64 {
	return &p.angularAcceleration
}

func (p *Physics) AngularDrag() *float64 {
	return &p.angularDrag
}

func (p *Physics) MaxAngular() *float64 {
	return &p.maxAngular
}

func (p *Physics) MaxVelocity() *math.GxlPoint {
	return &p.maxVelocity
}

func (p *Physics) Acceleration() *math.GxlPoint {
	return &p.acceleration
}

func (p *Physics) Drag() *math.GxlPoint {
	return &p.drag
}

func computeVelocity(velocity float64, acceleration float64, drag float64, max float64, elapsed float64) float64 {
	if acceleration != 0 {
		velocity += acceleration * elapsed
	} else if drag != 0 {
		drag := drag * elapsed
		if velocity-drag > 0 {
			velocity -= drag
		} else if (velocity + drag) < 0 {
			velocity += drag
		} else {
			velocity = 0
		}
	}
	if (velocity != 0) && (max != 0) {
		if velocity > max {
			velocity = max
		} else if velocity < -max {
			velocity = -max
		}
	}

	return velocity
}
