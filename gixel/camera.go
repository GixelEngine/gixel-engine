package gixel

import (
	"github.com/GixelEngine/gixel-engine/gixel/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type GxlCameraManager struct {
	state   GxlState
	cameras []*GxlCamera
}

func (cm *GxlCameraManager) Init(game *GxlGame) {
	cm.state = game.state
	cm.cameras = make([]*GxlCamera, 0)
	cm.New(0, 0, game.width, game.height)
}

func (cm *GxlCameraManager) New(x, y float64, w, h int) *GxlCamera {
	newCam := &GxlCamera{x: x, y: y, w: w, h: h}
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

func (cm *GxlCameraManager) Draw(screen *ebiten.Image) {
	cm.cameras[0].screen = screen

	cm.state.Range(func(idx int, obj GxlBasic) bool {
		if !*obj.Exists() || !*obj.Visible() {
			return true
		}

		obj.Draw()

		return true
	})

	if len(cm.cameras) == 0 {
		return
	}

	// TODO: Check if default camera has shaders

	for _, c := range cm.cameras[1:] {
		screen.DrawImage(c.screen, &ebiten.DrawImageOptions{})
		c.screen.Clear()
	}
}

type GxlCamera struct {
	BaseGxlBasic
	x, y   float64
	w, h   int
	scroll math.GxlPoint
	screen *ebiten.Image
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
