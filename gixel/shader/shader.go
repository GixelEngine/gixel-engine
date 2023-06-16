package shader

import "github.com/hajimehoshi/ebiten/v2"

type BaseGxlShader struct {
	key      string
	shader   *ebiten.Shader
	uniforms map[string]interface{}
}

func NewShader(key string) *BaseGxlShader {
	return &BaseGxlShader{
		key: key,
	}
}

func (s *BaseGxlShader) Init(shader *ebiten.Shader) {
	s.shader = shader
	s.uniforms = make(map[string]interface{})
}

func (s *BaseGxlShader) Key() string {
	return s.key
}

func (s *BaseGxlShader) Shader() *ebiten.Shader {
	return s.shader
}

func (s *BaseGxlShader) Uniforms() map[string]interface{} {
	return s.uniforms
}

func (s *BaseGxlShader) Update(elapsed float64) {
}

type GxlShader interface {
	Init(shader *ebiten.Shader)
	Key() string
	Shader() *ebiten.Shader
	Uniforms() map[string]interface{}
	Update(elapsed float64)
}
