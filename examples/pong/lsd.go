package main

import (
	"github.com/GixelEngine/gixel-engine/gixel/shader"
)

type LSDShader struct {
	shader.BaseGxlShader
}

func NewLSDShader(white bool) *LSDShader {
	val := 0
	if white {
		val = 1
	}
	return &LSDShader{
		BaseGxlShader: *shader.NewShader("assets/shaders/lsd.kage", map[string]interface{}{
			"Time":    float32(0),
			"IsWhite": float32(val),
		}),
	}
}

func (s *LSDShader) Update(elapsed float64) {
	s.Uniforms()["Time"] = s.Uniforms()["Time"].(float32) + float32(elapsed)
}
