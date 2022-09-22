package main

import (
	"image/color"

	"github.com/GixelEngine/gixel-engine/gixel"
)

type Letter struct {
	gixel.BaseGxlSprite
	alpha float64
	char  string
}

func NewLetter(x float64, y float64, char string) *Letter {
	l := &Letter{char: char}
	l.SetPosition(x, y)
	return l
}

func (l *Letter) Init(game *gixel.GxlGame) {
	l.BaseGxlSprite.Init(game)
	l.ApplyGraphic(l.Game().Graphics().Get(l.char))
	*l.Color() = color.RGBA{G: 255, A: 255}
	l.alpha = 255
}

func (l *Letter) Update(elapsed float64) error {
	err := l.BaseGxlSprite.Update(elapsed)
	if err != nil {
		return err
	}

	l.alpha -= elapsed * 300

	l.Color().A = uint8(l.alpha)

	if l.alpha <= 0 {
		l.Destroy()
	}

	return nil
}
