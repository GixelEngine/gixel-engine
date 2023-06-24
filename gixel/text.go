package gixel

import (
	"image/color"
	"log"

	"github.com/GixelEngine/gixel-engine/gixel/cache"
	"github.com/GixelEngine/gixel-engine/gixel/font"
	"github.com/GixelEngine/gixel-engine/gixel/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type BaseGxlText struct {
	BaseGxlSprite
	text       string
	fontPreset *font.GxlFontPreset
	needUpdate bool
	screenPos  *math.GxlPoint
}

func (t *BaseGxlText) Init(game *GxlGame) {
	t.BaseGxlSprite.Init(game)
	t.needUpdate = false
	t.updateGraphic()
}

// NewText creates a new instance of GxlText with a given font in a given position.
func NewText(x, y float64, text string, fontPreset *font.GxlFontPreset) GxlText {
	t := &BaseGxlText{}
	t.SetPosition(x, y)
	t.fontPreset = fontPreset
	t.text = text
	return t
}

func (t *BaseGxlText) SetText(text string) {
	t.text = text
	t.needUpdate = true
}

func (t *BaseGxlText) SetScreenPosition(pos *math.GxlPoint) {
	if pos == nil {
		log.Fatal("cannot set nil position")
	}
	t.screenPos = pos
}

func (t *BaseGxlText) SetFontPreset(fontPreset *font.GxlFontPreset) {
	if fontPreset == nil {
		log.Fatal("cannot set nil font")
	}
	t.fontPreset = fontPreset
	t.needUpdate = true
}

func (t *BaseGxlText) updateGraphic() {
	rect := text.BoundString(t.fontPreset.GetFace(), t.text)
	p := rect.Size()
	t.w, t.h = p.X, p.Y

	t.graphic = t.Game().Graphics().LoadGraphicFromImage(ebiten.NewImage(t.w, t.h), cache.CacheOptions{NoCache: true})
	text.Draw(t.graphic.GetFrame(0), t.text, t.fontPreset.GetFace(), -rect.Min.X, -rect.Min.Y, color.White)

	if t.screenPos == nil {
		return
	}

	t.SetPosition(float64(t.game.W())*t.screenPos.X-float64(t.w/2), float64(t.game.H())*t.screenPos.Y-float64(t.h/2))
	t.screenPos = nil
}

func (t *BaseGxlText) Draw() {
	if t.needUpdate {
		t.updateGraphic()
	}
	t.needUpdate = false

	t.BaseGxlSprite.Draw()
}

type GxlText interface {
	GxlSprite
	SetText(text string)
	SetFontPreset(fontPreset *font.GxlFontPreset)
	SetScreenPosition(pos *math.GxlPoint)
	updateGraphic()
}
