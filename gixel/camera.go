package gixel

import (
	"github.com/GixelEngine/gixel-engine/gixel/math"
	"github.com/GixelEngine/gixel-engine/gixel/shader"
	"github.com/hajimehoshi/ebiten/v2"
)

type GxlCameraManager struct {
	game       *GxlGame
	state      GxlState
	cameras    []*GxlCamera
	drawBuffer *ebiten.Image
}

func (cm *GxlCameraManager) Init(game *GxlGame) {
	cm.game = game
	cm.state = game.state
	cm.cameras = make([]*GxlCamera, 0)
	cm.New(0, 0, game.width, game.height)
	cm.drawBuffer = ebiten.NewImage(game.width, game.height)
}

func (cm *GxlCameraManager) New(x, y float64, w, h int) *GxlCamera {
	newCam := &GxlCamera{game: cm.game, x: x, y: y, w: w, h: h, filters: []shader.GxlShader{}}
	cm.cameras = append(cm.cameras, newCam)
	newCam.screen = ebiten.NewImage(w, h)

	return newCam
}

func (cm *GxlCameraManager) Get(idx int) *GxlCamera {
	if idx < 0 || idx >= len(cm.cameras) {
		return nil
	}

	return cm.cameras[idx]
}

func (cm *GxlCameraManager) GetDefault() *GxlCamera {
	return cm.cameras[0]
}

func (cm *GxlCameraManager) Update(elapsed float64) {
	for _, camera := range cm.cameras {
		camera.UpdateFilters(elapsed)
	}
}

func (cm *GxlCameraManager) Draw(screen *ebiten.Image) {
	cm.cameras[0].screen = screen

	hasFilters := len(cm.cameras[0].filters) > 0
	if hasFilters {
		cm.drawBuffer.Clear()
		cm.cameras[0].screen = cm.drawBuffer
	}

	cm.state.Range(func(idx int, obj GxlBasic) bool {
		if !*obj.Exists() || !*obj.Visible() {
			return true
		}

		obj.Draw()

		return true
	})

	if hasFilters {
		for i, filter := range cm.cameras[0].filters {
			screen.DrawRectShader(cm.game.width, cm.game.height, filter.Shader(), &ebiten.DrawRectShaderOptions{
				Uniforms: filter.Uniforms(),
				Images:   [4]*ebiten.Image{cm.drawBuffer, nil, nil, nil},
			})

			if i < len(cm.cameras[0].filters)-1 {
				cm.drawBuffer.DrawImage(screen, &ebiten.DrawImageOptions{})
			}
		}
	}

	if len(cm.cameras) == 0 {
		return
	}

	// TODO: take camera position into account
	// TODO: take filters into account
	for _, c := range cm.cameras[1:] {
		screen.DrawImage(c.screen, &ebiten.DrawImageOptions{})
		c.screen.Clear()
	}
}

type GxlCamera struct {
	game    *GxlGame
	x, y    float64
	w, h    int
	scroll  math.GxlPoint
	screen  *ebiten.Image
	filters []shader.GxlShader
}

func (c *GxlCamera) SetFilters(filters []shader.GxlShader) {
	c.filters = []shader.GxlShader{}
	for _, filter := range filters {
		filter.Init(c.game.Shaders().LoadShader(filter.Path()))
		c.filters = append(c.filters, filter)
	}
}

func (c *GxlCamera) ClearFilters() {
	c.filters = []shader.GxlShader{}
}

func (c *GxlCamera) UpdateFilters(elapsed float64) {
	for _, filter := range c.filters {
		filter.Update(elapsed)
	}
}

func (c *GxlCamera) X() *float64 {
	return &c.x
}

func (c *GxlCamera) Y() *float64 {
	return &c.y
}

func (c *GxlCamera) W() int {
	return c.w
}

func (c *GxlCamera) H() int {
	return c.h
}

func (c *GxlCamera) ContainsRect(rect *math.GxlRectangle) bool {
	return *rect.X()+*rect.W() > c.Scroll().X && *rect.X() < float64(c.w)+c.Scroll().X && *rect.Y()+*rect.H() > c.Scroll().Y && *rect.Y() < float64(c.h)+c.Scroll().Y
}

func (c *GxlCamera) Scroll() *math.GxlPoint {
	return &c.scroll
}

func (c *GxlCamera) Screen() *ebiten.Image {
	return c.screen
}
