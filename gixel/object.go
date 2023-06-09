package gixel

import (
	"github.com/GixelEngine/gixel-engine/gixel/math"
)

// TODO: Figure out a place to have this type
type GxlDirection uint16

const (
	None  GxlDirection = 0x0000
	Up    GxlDirection = 0x0001
	Down  GxlDirection = 0x0010
	Left  GxlDirection = 0x0100
	Right GxlDirection = 0x1000
)

// TODO: Consider removing/merging with GxlSprite
type BaseGxlObject struct {
	BaseGxlBasic
	x, y            float64
	w, h            int
	angle           float64
	angleMultiplier float64
	scale           *math.GxlPoint
	scaleMultiplier *math.GxlPoint
	scrollFactor    *math.GxlPoint
	offset          *math.GxlPoint
}

func (o *BaseGxlObject) Init(game *GxlGame) {
	o.BaseGxlBasic.Init(game)
	o.scale = math.NewPoint(1, 1)
	o.scaleMultiplier = math.NewPoint(1, 1)
	o.angleMultiplier = 1
	o.scrollFactor = math.NewPoint(1, 1)
	o.offset = math.NewPoint(0, 0)
}

func (o *BaseGxlObject) X() *float64 {
	return &o.x
}

func (o *BaseGxlObject) Y() *float64 {
	return &o.y
}

func (o *BaseGxlObject) W() *int {
	return &o.w
}

func (o *BaseGxlObject) H() *int {
	return &o.h
}

func (o *BaseGxlObject) GetPosition() *math.GxlPoint {
	return math.NewPoint(o.x, o.y)
}

func (o *BaseGxlObject) SetPosition(x, y float64) {
	o.x, o.y = x, y
}

func (o *BaseGxlObject) ScreenPosition() *math.GxlPoint {
	return math.NewPoint(o.x-o.camera.scroll.X*o.scrollFactor.X, o.y-o.camera.scroll.Y*o.scrollFactor.Y)
}

func (o *BaseGxlObject) GetSize() *math.GxlPoint {
	return math.NewPoint(float64(o.w), float64(o.h))
}

func (o *BaseGxlObject) SetSize(w, h int) {
	o.w, o.h = w, h
}

func (o *BaseGxlObject) Scale() *math.GxlPoint {
	return o.scale
}

func (o *BaseGxlObject) ScaleMultiplier() *math.GxlPoint {
	return o.scaleMultiplier
}

func (o *BaseGxlObject) ScrollFactor() *math.GxlPoint {
	return o.scrollFactor
}

func (o *BaseGxlObject) Angle() *float64 {
	return &o.angle
}

func (o *BaseGxlObject) AngleMultiplier() *float64 {
	return &o.angleMultiplier
}

func (o *BaseGxlObject) Offset() *math.GxlPoint {
	return o.offset
}

func (o *BaseGxlObject) Bounds() *math.GxlRectangle {
	return math.NewRectangle(o.x+o.offset.X, o.y+o.offset.Y, float64(o.w), float64(o.h))
}

func (o *BaseGxlObject) SetHitbox(x, y float64, w, h int) {
	o.offset.X, o.offset.Y = x, y
	o.SetSize(w, h)
}

func (o *BaseGxlObject) Overlaps(obj GxlObject) bool {
	if !o.exists || !*obj.Exists() {
		return false
	}

	return o.Bounds().Overlaps(obj.Bounds())
}

func (o *BaseGxlObject) OnScreen() bool {
	return o.camera.ContainsRect(o.Bounds())
}

type GxlObject interface {
	GxlBasic
	X() *float64
	Y() *float64
	W() *int
	H() *int
	GetPosition() *math.GxlPoint
	SetPosition(x, y float64)
	ScreenPosition() *math.GxlPoint
	GetSize() *math.GxlPoint
	SetSize(w, h int)
	Scale() *math.GxlPoint
	ScaleMultiplier() *math.GxlPoint
	ScrollFactor() *math.GxlPoint
	Angle() *float64
	AngleMultiplier() *float64
	Offset() *math.GxlPoint
	Bounds() *math.GxlRectangle
	SetHitbox(x, y float64, w, h int)
	Overlaps(obj GxlObject) bool
	OnScreen() bool
}
