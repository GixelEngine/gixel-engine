package graphic

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type cacheRecord struct {
	persist bool
	graphic *GxlGraphic
}

type GxlGraphicCache struct {
	cache map[string]*cacheRecord
	fs    *embed.FS
}

func NewGraphicCache(fs *embed.FS) *GxlGraphicCache {
	return &GxlGraphicCache{
		cache: make(map[string]*cacheRecord),
		fs:    fs,
	}
}

func (gc *GxlGraphicCache) Add(graphic *GxlGraphic, key string, persist bool) *GxlGraphic {
	if graphic == nil || key == "" {
		return nil
	}

	gc.cache[key] = &cacheRecord{persist: persist, graphic: graphic}

	return graphic
}

func (gc *GxlGraphicCache) Get(key string) *GxlGraphic {
	record, ok := gc.cache[key]
	if !ok || record == nil {
		return nil
	}

	return record.graphic
}

func (gc *GxlGraphicCache) Clear() {
	for _, record := range gc.cache {
		if !record.persist {
			record = nil
		}
	}
}

// MakeGraphic creates a new ebiten.Image fills it with a given color.
//
// Returns a pointer of the GxlGraphic.
func (gc *GxlGraphicCache) MakeGraphic(w, h int, c color.Color, opt CacheOptions) *GxlGraphic {
	if opt.Key == "" {
		r, g, b, a := c.RGBA()
		opt.Key = fmt.Sprintf("%d_%d_%02x%02x%02x%02x", w, h, r, g, b, a)
	}

	if !opt.Unique {
		g := gc.Get(opt.Key)
		if g != nil {
			return g
		}
	} else {
		opt.Key += "_" + uuid.New().String()
	}

	img := ebiten.NewImage(w, h)
	img.Fill(c)

	return gc.Add(
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
func (gc *GxlGraphicCache) LoadGraphic(path string, opt CacheOptions) *GxlGraphic {
	if opt.Key == "" {
		opt.Key = path
	}

	if !opt.Unique {
		g := gc.Get(opt.Key)
		if g != nil {
			return g
		}
	} else {
		opt.Key += "_" + uuid.New().String()
	}

	img, _, err := ebitenutil.NewImageFromFileSystem(gc.fs, path)
	if err != nil {
		log.Panicln(err)
		// TODO: Error handling, default value?
	}

	return gc.Add(
		&GxlGraphic{
			frames: []*ebiten.Image{img},
		},
		opt.Key,
		opt.Persist,
	)
}

func (gc *GxlGraphicCache) LoadGraphicFromImage(img *ebiten.Image, opt CacheOptions) *GxlGraphic {
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
		g := gc.Get(opt.Key)
		if g != nil {
			return g
		}
	} else {
		opt.Key += "_" + uuid.New().String()
	}

	return gc.Add(g, opt.Key, opt.Persist)
}

// MakeGraphic creates a new ebiten.Image from a file path.
//
// Returns a pointer of the GxlGraphic.
func (gc *GxlGraphicCache) LoadAnimatedGraphic(path string, fw, fh int, opt CacheOptions) *GxlGraphic {
	if opt.Key == "" {
		opt.Key = path
	}

	if !opt.Unique {
		g := gc.Get(opt.Key)
		if g != nil {
			return g
		}
	} else {
		opt.Key += "_" + uuid.New().String()
	}

	img, _, err := ebitenutil.NewImageFromFileSystem(gc.fs, path)
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

	return gc.Add(
		&GxlGraphic{
			frames: frames,
		},
		opt.Key,
		opt.Persist,
	)
}

type CacheOptions struct {
	Key     string
	Unique  bool
	Persist bool
	NoCache bool
}
