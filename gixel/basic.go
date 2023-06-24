package gixel

type BaseGxlBasic struct {
	game    *GxlGame
	id      int
	visible bool
	exists  bool
	camera  *GxlCamera
	// TODO: Do we want current state ref here?
}

func (b *BaseGxlBasic) Init(g *GxlGame) {
	b.game = g
	b.id = g.GenerateID()
	b.visible = true
	b.exists = true
	if b.camera == nil {
		b.camera = g.state.Cameras().GetDefault()
	}
	b.game.Debug().Counters.InitCount.Increment()
}

func (b *BaseGxlBasic) SwitchCamera(camera *GxlCamera) {
	b.camera = camera
}

func (b *BaseGxlBasic) Camera() *GxlCamera {
	return b.camera
}

func (b *BaseGxlBasic) Destroy() {
	b.exists = false
}

func (b *BaseGxlBasic) Draw() {
	b.game.Debug().Counters.DrawCount.Increment()
}

func (b *BaseGxlBasic) Update(elapsed float64) error {
	b.game.Debug().Counters.UpdateCount.Increment()
	return nil
}

func (b *BaseGxlBasic) GetID() int {
	return b.id
}

func (b *BaseGxlBasic) Visible() *bool {
	return &b.visible
}

func (b *BaseGxlBasic) Exists() *bool {
	return &b.exists
}

func (b *BaseGxlBasic) Game() *GxlGame {
	return b.game
}

type GxlBasic interface {
	Init(g *GxlGame)
	Destroy()
	Draw()
	Update(elapsed float64) error
	GetID() int
	Visible() *bool
	Exists() *bool
	Game() *GxlGame
	Camera() *GxlCamera
	SwitchCamera(camera *GxlCamera)
}
