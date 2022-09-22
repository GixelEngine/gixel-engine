package states

import (
	"github.com/GixelEngine/gixel-engine/examples/testing/entities"
	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/systems/collision"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayState struct {
	gixel.BaseGxlState
	player *entities.Player
	walls  gixel.GxlGroup
	wall   *entities.Wall
	//Systems
	collision.Collision
}

func (s *PlayState) Init(game *gixel.GxlGame) {
	s.BaseGxlState.Init(game)

	s.Collision.Init(s)

	s.walls = gixel.NewGroup(0)
	s.Add(s.walls)
	for i := 0.0; i < 4; i++ {
		s.walls.Recycle(func() gixel.GxlBasic {
			return entities.NewWall(float64(100+32*i), 150+2*i)
		})
	}

	s.wall = entities.NewWall(25, 200)
	s.Add(s.wall)
	s.wall.Color().R = 255

	s.player = entities.NewPlayer(100, 100)
	s.Add(s.player)
}

func (s *PlayState) Update(elapsed float64) error {
	err := s.BaseGxlState.Update(elapsed)
	if err != nil {
		return err
	}

	s.CollideObjects(s.player, s.wall)
	s.CollideObjectGroup(s.player, s.walls)
	s.CollideGroups(s.walls, s.walls)

	s.Collision.Update(elapsed)

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		fullscreen := ebiten.IsFullscreen()
		fullscreen = !fullscreen
		ebiten.SetFullscreen(fullscreen)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		*s.Game().Timescale() = 0.5
	}

	return nil
}
