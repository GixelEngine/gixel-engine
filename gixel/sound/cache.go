package sound

import (
	"embed"
	"io"
	"log"

	"github.com/GixelEngine/gixel-engine/gixel/cache"
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

type GxlSoundCache struct {
	cache *cache.GxlCache[GxlSoundAsset]
}

func NewSoundCache(fs *embed.FS) *GxlSoundCache {
	return &GxlSoundCache{cache: cache.NewCache[GxlSoundAsset](fs)}
}

func (gsc *GxlSoundCache) Add(data *GxlSoundAsset, key string, persist bool) {
	gsc.cache.Add(data, key, persist)
}

func (gsc *GxlSoundCache) Get(key string) *GxlSoundAsset {
	return gsc.cache.Get(key)
}

func (gsc *GxlSoundCache) Clear() {
	gsc.cache.Clear()
}

// LoadSound creates a new GxlSoundAsset from a file path.
//
// Returns a pointer of the GxlSoundAsset.
// TODO: Add file types support
func (gsc *GxlSoundCache) LoadSound(path string, opt cache.CacheOptions) *GxlSoundAsset {
	if opt.Key == "" {
		opt.Key = path
	}

	if !opt.Unique {
		g := gsc.cache.Get(opt.Key)
		if g != nil {
			return g
		}
	} else {
		opt.Key += "_" + uuid.New().String()
	}

	file, err := gsc.cache.FS.Open(path)
	if err != nil {
		log.Panicln(err)
		// TODO: Error handling, default value?
	}

	s, err := vorbis.DecodeWithoutResampling(file)
	if err != nil {
		log.Panicln(err)
	}

	bs := make([]byte, s.Length())
	_, err = s.Read(bs)
	if err != nil && err != io.EOF {
		log.Panicln(err)
	}

	return gsc.cache.Add(
		&GxlSoundAsset{data: bs},
		opt.Key,
		opt.Persist,
	)
}
