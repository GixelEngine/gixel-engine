package shader

import (
	"bufio"
	"embed"
	"io"
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

func (gsc *GxlShaderCache) Register(path string) *ebiten.Shader {
	file, err := gsc.cache.FS.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		log.Panicln(err)
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		log.Panicln(err)
	}

	s, err := ebiten.NewShader(bs)
	if err != nil {
		log.Panicln(err)
	}
	gsc.cache.Add(s, path, false)

	return s
}

func (gsc *GxlShaderCache) LoadShader(path string) *ebiten.Shader {
	s := gsc.cache.Get(path)
	if s != nil {
		return s
	}

	return gsc.Register(path)
}
