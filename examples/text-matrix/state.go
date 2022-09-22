package main

import (
	"math/rand"
	"time"

	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/font"
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayState struct {
	gixel.BaseGxlState
	letters    gixel.GxlGroup
	spawnTimer *gixel.GxlTimer
}

func (s *PlayState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)

	s.letters = gixel.NewGroup(6000)
	s.Add(s.letters)

	f := font.NewFont("assets/Boku2-Bold.otf")

	var katakana []rune

	alphabet := `アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲンガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ`

	for _, letter := range alphabet {
		katakana = append(katakana, letter)
		s := gixel.NewText(0, 0, string(letter), f.GetPreset(32))
		s.Init(game)
		game.Graphics().Add(s.Graphic(), string(letter), false)
	}

	rand := rand.New(rand.NewSource(time.Now().Unix()))
	s.spawnTimer = gixel.NewLoopTimer(1.0 / 60).SetCallback(func(totalElapsed float64, iteration int) {
		for i := 0; i < 50; i++ {
			char := string(katakana[rand.Int()%(len(katakana)-1)])
			letter := s.letters.Recycle(func() gixel.GxlBasic {
				return NewLetter(0, 0, char)
			}).(*Letter)

			*letter.X() = rand.Float64() * GAME_WIDTH
			*letter.Y() = rand.Float64() * GAME_HEIGHT
			letter.alpha = 255
			*letter.Exists() = true
		}
	})
	s.spawnTimer.Start()
}

func (s *PlayState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	s.spawnTimer.Update(elapsed)

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.Cameras().GetDefault().Scroll().X += 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.Cameras().GetDefault().Scroll().X -= 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.Cameras().GetDefault().Scroll().Y -= 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s.Cameras().GetDefault().Scroll().Y += 10
	}

	return nil
}
