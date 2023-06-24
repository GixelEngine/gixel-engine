package main

import (
	"github.com/GixelEngine/gixel-engine/gixel/shader"
)

type CRTShader struct {
	shader.BaseGxlShader
}

func NewCRTShader() *CRTShader {
	return &CRTShader{
		BaseGxlShader: *shader.NewShader("assets/shaders/crt.kage", map[string]interface{}{}),
	}
}

func (s *CRTShader) Update(elapsed float64) {
}
