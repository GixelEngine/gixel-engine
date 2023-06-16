package shader

import (
	"embed"
	"log"

	"github.com/GixelEngine/gixel-engine/gixel/cache"
	"github.com/hajimehoshi/ebiten/v2"
)

type GxlShaderCache struct {
	cache *cache.GxlCache[ebiten.Shader]
}

func NewShaderCache(fs *embed.FS) *GxlShaderCache {
	return &GxlShaderCache{cache: cache.NewCache[ebiten.Shader](fs)}
}

func (gsc *GxlShaderCache) Register(src []byte, key string) {
	s, err := ebiten.NewShader(src)
	if err != nil {
		log.Panicln(err)
	}
	gsc.cache.Add(s, key, false)
}

func (gsc *GxlShaderCache) Get(key string) *ebiten.Shader {
	return gsc.cache.Get(key)
}

func (gsc *GxlShaderCache) Clear() {
	gsc.cache.Clear()
}
