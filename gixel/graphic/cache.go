package graphic

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/GixelEngine/gixel-engine/gixel/cache"
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GxlGraphicCache struct {
	cache *cache.GxlCache[GxlGraphic]
}

func NewGraphicCache(fs *embed.FS) *GxlGraphicCache {
	return &GxlGraphicCache{cache: cache.NewCache[GxlGraphic](fs)}
}

func (gc *GxlGraphicCache) Add(data *GxlGraphic, key string, persist bool) {
	gc.cache.Add(data, key, persist)
}

func (gc *GxlGraphicCache) Get(key string) *GxlGraphic {
	return gc.cache.Get(key)
}

func (gc *GxlGraphicCache) Clear() {
	gc.cache.Clear()
}

// MakeGraphic creates a new ebiten.Image fills it with a given color.
//
// Returns a pointer of the GxlGraphic.
func (gc *GxlGraphicCache) MakeGraphic(w, h int, c color.Color, opt cache.CacheOptions) *GxlGraphic {
	if opt.Key == "" {
		r, g, b, a := c.RGBA()
		opt.Key = fmt.Sprintf("%d_%d_%02x%02x%02x%02x", w, h, r, g, b, a)
	}

	if !opt.Unique {
		g := gc.cache.Get(opt.Key)
		if g != nil {
			return g
		}
	} else {
		opt.Key += "_" + uuid.New().String()
	}

	img := ebiten.NewImage(w, h)
	img.Fill(c)

	return gc.cache.Add(
		&GxlGraphic{
			frames: []*ebiten.Image{img},
		},
		opt.Key,
		opt.Persist,
	)
}

// MakeGraphic creates a new ebiten.Image from a file path.
//
// Returns a pointer of the GxlGraphic.
func (gc *GxlGraphicCache) LoadGraphic(path string, opt cache.CacheOptions) *GxlGraphic {
	if opt.Key == "" {
		opt.Key = path
	}

	if !opt.Unique {
		g := gc.cache.Get(opt.Key)
		if g != nil {
			return g
		}
	} else {
		opt.Key += "_" + uuid.New().String()
	}

	img, _, err := ebitenutil.NewImageFromFileSystem(gc.cache.FS, path)
	if err != nil {
		log.Panicln(err)
		// TODO: Error handling, default value?
	}

	return gc.cache.Add(
		&GxlGraphic{
			frames: []*ebiten.Image{img},
		},
		opt.Key,
		opt.Persist,
	)
}

func (gc *GxlGraphicCache) LoadGraphicFromImage(img *ebiten.Image, opt cache.CacheOptions) *GxlGraphic {
	g := &GxlGraphic{
		frames: []*ebiten.Image{img},
	}
	if opt.NoCache {
		return g
	}

	if opt.Key == "" {
		opt.Unique = true
	}

	if !opt.Unique {
		g := gc.cache.Get(opt.Key)
		if g != nil {
			return g
		}
	} else {
		opt.Key += "_" + uuid.New().String()
	}

	return gc.cache.Add(g, opt.Key, opt.Persist)
}

// MakeGraphic creates a new ebiten.Image from a file path.
//
// Returns a pointer of the GxlGraphic.
func (gc *GxlGraphicCache) LoadAnimatedGraphic(path string, fw, fh int, opt cache.CacheOptions) *GxlGraphic {
	if opt.Key == "" {
		opt.Key = path
	}

	if !opt.Unique {
		g := gc.cache.Get(opt.Key)
		if g != nil {
			return g
		}
	} else {
		opt.Key += "_" + uuid.New().String()
	}

	img, _, err := ebitenutil.NewImageFromFileSystem(gc.cache.FS, path)
	if err != nil {
		log.Panicln(err)
		// TODO: Error handling, default value?
	}

	w, h := img.Size()
	frameRows := (h / fh)
	frameCols := (w / fw)

	frames := make([]*ebiten.Image, 0, frameCols*frameRows)

	for i := 0; i < frameRows; i++ {
		for j := 0; j < frameCols; j++ {
			frameRect := image.Rect(j*fw, i*fh, j*fw+fw, i*fh+fh)
			frames = append(frames, img.SubImage(frameRect).(*ebiten.Image))
		}
	}

	return gc.cache.Add(
		&GxlGraphic{
			frames: frames,
		},
		opt.Key,
		opt.Persist,
	)
}
