package main

import (
	"fmt"

	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/font"
	"github.com/GixelEngine/gixel-engine/gixel/math"
	"github.com/GixelEngine/gixel-engine/gixel/shader"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const DISTANCE_FROM_WALL = 64

type PlayState struct {
	gixel.BaseGxlState
	paddle1, paddle2 *Paddle
	score1, score2   int
	scoreText        gixel.GxlText
	paddles          gixel.GxlGroup
	ball             *Ball
	startTimer       *gixel.GxlTimer
	timerText        gixel.GxlText
}

func (s *PlayState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)

	s.Camera().SetFilters([]shader.GxlShader{
		NewLSDShader(true),
		NewLSDShader(false),
		NewCRTShader(),
	})

	s.paddles = gixel.NewGroup(0)
	s.Add(s.paddles)

	s.paddle1 = NewPaddle(DISTANCE_FROM_WALL, GAME_HEIGHT/2)
	s.paddles.Add(s.paddle1)
	s.paddle1.SetControls(ebiten.KeyW, ebiten.KeyS)

	s.paddle2 = NewPaddle(GAME_WIDTH-DISTANCE_FROM_WALL-float64(*s.paddle1.W()), GAME_HEIGHT/2)
	s.paddles.Add(s.paddle2)
	s.paddle2.SetControls(ebiten.KeyUp, ebiten.KeyDown)

	s.ball = NewBall()
	s.Add(s.ball)

	f := font.NewFont("assets/nokiafc22.ttf")
	s.scoreText = gixel.NewText(0, 0, "0 - 0", f.GetPreset(64))

	s.Add(s.scoreText)

	s.timerText = gixel.NewText(0, 0, "0", f.GetPreset(128))
	*s.timerText.Visible() = false
	countdown := []string{"2", "1", "START"}
	s.Add(s.timerText)

	s.startTimer = gixel.NewIterationTimer(4, 0.5).
		SetCallback(func(totalElapsed float64, iteration int) {
			if iteration-1 < len(countdown) {
				s.timerText.SetText(countdown[iteration-1])
				s.timerText.SetScreenPosition(math.NewPoint(0.5, 0.5))
			} else {
				*s.timerText.Visible() = false
				*(*s.ball).Visible() = true
				*(*s.ball).Active() = true
			}
		}).
		SetOnStart(func(totalElapsed float64, iteration int) {
			s.ball.Spawn()
			*(*s.ball).Visible() = false
			*(*s.ball).Active() = false
			*s.timerText.Visible() = true
			s.timerText.SetText("3")
			s.timerText.SetScreenPosition(math.NewPoint(0.5, 0.5))

		})
	s.startTimer.Start()

	s.Game().SoundManager().PlayMusic("assets/sounds/bg_music.ogg", 0.5, true)
}

func (s *PlayState) Update(elapsed float64) error {
	*s.Game().Timescale() = 1

	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	s.startTimer.Update(elapsed)

	if *s.ball.X() > GAME_WIDTH-float64(*s.ball.W()) {
		s.startTimer.Restart()
		s.score2++
	} else if *s.ball.X() < 0 {
		s.startTimer.Restart()
		s.score1++
	}

	checkCollision := true
	ballX := *s.ball.X()
	// Check if ball is in dead zone
	if *s.ball.Visible() && (ballX < DISTANCE_FROM_WALL+35 || ballX+float64(*s.ball.W()) > GAME_WIDTH-DISTANCE_FROM_WALL-35) {
		*s.Game().Timescale() = 0.25
		checkCollision = false
	}

	// Check collision only outside of dead zone
	if checkCollision {
		s.OverlapsObjectGroup(s.ball, s.paddles, func(obj1, obj2 *gixel.GxlObject) {
			dir := 1.0
			if *(*obj1).X() < GAME_WIDTH/2 {
				*(*obj1).X() = DISTANCE_FROM_WALL + 48 + 1
			} else {
				*(*obj1).X() = GAME_WIDTH - DISTANCE_FROM_WALL - 48 - 32 - 1
				dir = -1
			}
			halfPaddle := float64(*(*obj2).H() / 2)
			relativeY := (*(*obj2).Y() + halfPaddle) - (*(*obj1).Y())
			s.ball.FlipHorizontal(relativeY/halfPaddle, dir)
		})
	}

	s.scoreText.SetText(fmt.Sprintf("%d - %d", s.score1, s.score2))
	s.scoreText.SetScreenPosition(math.NewPoint(0.5, 0.1))

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		fullscreen := ebiten.IsFullscreen()
		fullscreen = !fullscreen
		ebiten.SetFullscreen(fullscreen)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		if s.Game().SoundManager().IsMusicPaused() {
			s.Game().SoundManager().ResumeMusic()
		} else {
			s.Game().SoundManager().PauseMusic()
		}
	}

	return nil
}
