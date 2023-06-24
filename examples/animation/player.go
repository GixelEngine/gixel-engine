package main

import (
	"fmt"

	"github.com/GixelEngine/gixel-engine/gixel"
	"github.com/GixelEngine/gixel-engine/gixel/cache"
	"github.com/GixelEngine/gixel-engine/gixel/systems/animation"
	"github.com/GixelEngine/gixel-engine/gixel/systems/flipping"
	"github.com/GixelEngine/gixel-engine/gixel/systems/physics"
)

type Player struct {
	gixel.BaseGxlSprite
	// Systems
	flipping.Flipping
	physics.Physics
	animation.Animation
	Movement
}

func NewPlayer(x, y float64) *Player {
	p := &Player{}
	p.SetPosition(x, y)

	return p
}

func (p *Player) Init(game *gixel.GxlGame) {
	p.BaseGxlSprite.Init(game)

	p.Flipping.Init(p)
	p.Physics.Init(p)
	p.Animation.Init(p)
	p.Movement.Init(p)

	p.ApplyGraphic(game.Graphics().LoadAnimatedGraphic("assets/player.png", 32, 32, cache.CacheOptions{}))
	p.SetHitbox(9, 14, 14, 18)
	p.SetFacingFlip(gixel.Right, false, false)
	p.SetFacingFlip(gixel.Left, true, false)

	p.AddAnimation("WalkFront", []int{0, 1, 0, 2}, 7, true)
	p.AddAnimation("WalkBack", []int{3, 4, 3, 5}, 7, true)
	p.AddAnimation("WalkSide", []int{6, 7, 6, 8}, 7, true)

	p.AddAnimation("StandFront", []int{0, 0, 9}, 5, false).
		SetCallback(0, func() { fmt.Println("yay") }).
		SetCallback(1, func() { fmt.Println("yay2") }).
		SetOnFinished(func() { fmt.Println("Wat") })
	p.AddAnimation("StandBack", []int{3, 3, 10}, 5, true)

}

func (p *Player) Update(elapsed float64) error {
	err := p.BaseGxlSprite.Update(elapsed)
	if err != nil {
		return err
	}

	p.Flipping.Update()
	p.Physics.Update(elapsed)
	p.Animation.Update(elapsed)
	p.Movement.Update(elapsed)

	p.PlayAnimation(p.Movement.GetAnimName(), false)
	return nil
}
