package main

import (
	_ "embed"

	"github.com/GixelEngine/gixel-engine/gixel/shader"
	"github.com/hajimehoshi/ebiten/v2"
)

type TestShader struct {
	shader.BaseGxlShader
	sum float64
}

//go:embed assets/shaders/shader.kage
var shaderProgram []byte

func NewTestShader(img *ebiten.Image) *TestShader {
	return &TestShader{
		BaseGxlShader: *shader.NewShader(shaderProgram, &ebiten.DrawRectShaderOptions{
			Uniforms: map[string]interface{}{
				"Idx": float32(0),
			},
			Images: [4]*ebiten.Image{img, nil, nil, nil}}),
		sum: 0,
	}
}

func (s *TestShader) Update(elapsed float64) {
	s.sum += elapsed
	s.Opts().Uniforms["Idx"] = float32(int(s.sum) % 3)
}
