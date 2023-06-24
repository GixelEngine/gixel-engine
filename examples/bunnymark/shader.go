package main

import (
	_ "embed"

	"github.com/GixelEngine/gixel-engine/gixel/shader"
)

type ColorSwapShader struct {
	shader.BaseGxlShader
	interval float64
	sum      float64
	idx      int
}

func NewColorSwapShader(interval float64) *ColorSwapShader {
	return &ColorSwapShader{
		BaseGxlShader: *shader.NewShader("assets/shaders/color_swap.kage", map[string]interface{}{}),
		interval:      interval,
		sum:           0,
		idx:           0,
	}
}

func (s *ColorSwapShader) Update(elapsed float64) {
	s.sum += elapsed
	if s.sum > s.interval {
		s.idx++
		s.idx %= 3
		s.sum -= s.interval
	}
	s.Uniforms()["Idx"] = float32(s.idx)
}
