package shader

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type BaseGxlShader struct {
	shader *ebiten.Shader
	opts   *ebiten.DrawRectShaderOptions
}

func NewShader(src []byte, opts *ebiten.DrawRectShaderOptions) *BaseGxlShader {
	s, err := ebiten.NewShader(src)
	if err != nil {
		log.Panicln(err)
	}

	return &BaseGxlShader{
		shader: s,
		opts:   opts,
	}
}

func (s *BaseGxlShader) Shader() *ebiten.Shader {
	return s.shader
}

func (s *BaseGxlShader) Opts() *ebiten.DrawRectShaderOptions {
	return s.opts
}

func (s *BaseGxlShader) Update(elapsed float64) {
}

type GxlShader interface {
	Shader() *ebiten.Shader
	Opts() *ebiten.DrawRectShaderOptions
	Update(elapsed float64)
}
