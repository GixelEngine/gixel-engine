package states

import (
	"image/color"
	"math"

	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/font"
	"github.com/GixelEngine/gixel-engine/gixel/graphic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

type MenuState struct {
	gixel.BaseGxlState
	text              gixel.GxlText
	textScaleSequence *gween.Sequence
	textAngleSequence *gween.Sequence
	textColorSequence *gween.Sequence
}

func (s *MenuState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)

	box1 := gixel.NewSprite(100, 100)
	box1.ApplyGraphic(game.Graphics().MakeGraphic(100, 100, color.RGBA{R: 255, A: 255}, graphic.CacheOptions{}))

	box2 := gixel.NewSprite(200, 200)
	box2.ApplyGraphic(game.Graphics().MakeGraphic(100, 100, color.RGBA{B: 255, A: 255}, graphic.CacheOptions{}))

	flixelFont := font.NewFont("assets/nokiafc22.ttf")
	fontSmall := flixelFont.GetPreset(8)
	fontMedium := flixelFont.GetPreset(16)
	fontLarge := flixelFont.GetPreset(32)

	textS := gixel.NewText(10, 10, "Hello World", fontSmall)
	s.Add(textS)
	textM := gixel.NewText(10, 20, "Hello World", fontMedium)
	s.Add(textM)
	textL := gixel.NewText(10, 40, "Hello World", fontLarge)
	s.Add(textL)

	text := gixel.NewText(100, 100, "Hello World", fontSmall)
	s.text = text

	s.textScaleSequence = gween.NewSequence(
		gween.New(1, 2, 5, ease.InQuad),
		gween.New(2, 1, 5, ease.InQuad),
	)
	s.textScaleSequence.SetLoop(-1)

	s.textAngleSequence = gween.NewSequence(
		gween.New(-math.Pi*0.25, math.Pi*0.25, 2, ease.InOutElastic),
		gween.New(math.Pi*0.25, -math.Pi*0.25, 2, ease.InOutElastic),
	)
	s.textAngleSequence.SetLoop(-1)

	s.textColorSequence = gween.NewSequence(
		gween.New(0, 255, 2, ease.InCubic),
		gween.New(255, 0, 2, ease.InCubic),
	)
	s.textColorSequence.SetLoop(-1)

	s.Add(box1)
	s.Add(box2)
	s.Add(text)
}

func (s *MenuState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	currentScale, _, _ := s.textScaleSequence.Update(float32(elapsed))
	s.text.Scale().X = float64(currentScale)
	s.text.Scale().Y = float64(currentScale)

	currentAngle, _, _ := s.textAngleSequence.Update(float32(elapsed))
	*s.text.Angle() = float64(currentAngle)

	currentRed, _, _ := s.textColorSequence.Update(float32(elapsed))
	s.text.Color().R = 255 - uint8(currentRed)
	s.text.Color().G = uint8(currentRed)

	if ebiten.IsKeyPressed(ebiten.KeyP) {
		s.Game().SwitchState(&PlayState{})
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		fullscreen := ebiten.IsFullscreen()
		fullscreen = !fullscreen
		ebiten.SetFullscreen(fullscreen)
	}

	return nil
}
