package gixel

import (
	"image/color"

	"github.com/GixelEngine/gixel-engine/gixel/graphic"
	"github.com/GixelEngine/gixel-engine/gixel/math"
	"github.com/GixelEngine/gixel-engine/gixel/shader"
	"github.com/hajimehoshi/ebiten/v2"
)

type BaseGxlSprite struct {
	BaseGxlObject
	graphic  *graphic.GxlGraphic
	frameIdx int
	color    color.RGBA // TODO: Think of a better name
	drawOpts *ebiten.DrawImageOptions
	shader   shader.GxlShader
	geom     ebiten.GeoM
}

func (s *BaseGxlSprite) Init(game *GxlGame) {
	s.BaseGxlObject.Init(game)
	s.drawOpts = &ebiten.DrawImageOptions{}
	s.color = color.RGBA{255, 255, 255, 255}
}

// NewSprite creates a new instance of GxlSprite in a given position.
func NewSprite(x, y float64) GxlSprite {
	s := &BaseGxlSprite{}
	s.SetPosition(x, y)
	return s
}

func (s *BaseGxlSprite) ApplyGraphic(graphic *graphic.GxlGraphic) {
	s.graphic = graphic
	s.SetSize(graphic.Size())
}

func (s *BaseGxlSprite) ApplyShader(shader shader.GxlShader) {
	shader.Init(s.Game().Shaders().Get(shader.Key()))
	s.shader = shader
}

func (s *BaseGxlSprite) Shader() shader.GxlShader {
	return s.shader
}

func (s *BaseGxlSprite) Graphic() *graphic.GxlGraphic {
	return s.graphic
}

func (s *BaseGxlSprite) FrameIdx() *int {
	return &s.frameIdx
}

func (s *BaseGxlSprite) Color() *color.RGBA {
	return &s.color
}

func (s *BaseGxlSprite) Update(elapsed float64) error {
	err := s.BaseGxlObject.Update(elapsed)
	if err != nil {
		return err
	}

	return nil
}

func (s *BaseGxlSprite) Draw() {
	if !s.OnScreen() {
		return
	}

	s.BaseGxlObject.Draw()
	if s.graphic == nil {
		return
	}

	s.geom.Reset()
	sxy := s.ScreenPosition()
	w, h := s.graphic.Size()
	s.geom.Translate(float64(-w/2), float64(-h/2))
	s.geom.Rotate(s.angle * s.angleMultiplier)
	s.geom.Scale(s.scale.X*s.scaleMultiplier.X, s.scale.Y*s.scaleMultiplier.Y)
	s.geom.Translate(float64(w/2), float64(h/2))
	s.geom.Translate(sxy.X, sxy.Y)

	if s.shader != nil && s.shader.Shader() != nil {
		s.camera.Screen().DrawRectShader(s.w, s.h, s.shader.Shader(), &ebiten.DrawRectShaderOptions{
			Uniforms: s.shader.Uniforms(),
			Images:   [4]*ebiten.Image{s.graphic.GetFrame(s.frameIdx), nil, nil, nil},
			GeoM:     s.geom,
		},
		)
	} else {
		s.drawOpts.GeoM = s.geom
		// // TODO: Add color for tinting/etc
		s.drawOpts.ColorM.Reset()
		s.drawOpts.ColorM.ScaleWithColor(s.color)
		s.camera.Screen().DrawImage(s.graphic.GetFrame(s.frameIdx), s.drawOpts)
	}
	// TODO: This currently prevents draw call batching, consider drawing in a separate run
	s.game.Debug().Collision.DrawBounds(s.camera.Screen(), *math.NewRectangle(sxy.X+s.offset.X, sxy.Y+s.offset.Y, float64(s.w), float64(s.h)))
}

type GxlSprite interface {
	GxlObject
	ApplyGraphic(graphic *graphic.GxlGraphic)
	ApplyShader(shader shader.GxlShader)
	Shader() shader.GxlShader
	Graphic() *graphic.GxlGraphic
	FrameIdx() *int
	Color() *color.RGBA
}
