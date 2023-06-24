package main

import (
	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayState struct {
	gixel.BaseGxlState
	gophers gixel.GxlGroup
}

func (s *PlayState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)

	s.gophers = gixel.NewGroup(0)
	s.Add(s.gophers)
}

func (s *PlayState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.gophers.Add(NewGopher(0, 0))
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		for i := 0; i <= 100; i++ {
			s.gophers.Add(NewGopher(0, 0))
		}
	}
	return nil
}
