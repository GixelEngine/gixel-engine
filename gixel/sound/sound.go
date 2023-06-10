package sound

import (
	"embed"

	"github.com/GixelEngine/gixel-engine/gixel/cache"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const sampleRate = 48000

type GxlSoundAsset struct {
	data []byte
}

type GxlSound struct {
	player *audio.Player
}

func (gs *GxlSound) Play() {
	if gs.player.IsPlaying() {
		return
	}

	gs.player.Rewind()
	gs.player.Play()
}

type GxlSoundManager struct {
	context     *audio.Context
	musicPlayer *audio.Player
	cache       *GxlSoundCache
}

func NewSoundManager(fs *embed.FS) *GxlSoundManager {
	return &GxlSoundManager{
		context: audio.NewContext(sampleRate),
		cache:   NewSoundCache(fs),
	}
}

func (gsm *GxlSoundManager) PlayMusic(file string) {
	sound := gsm.cache.LoadSound(file, cache.CacheOptions{})

	gsm.musicPlayer = gsm.context.NewPlayerFromBytes(sound.data)

	gsm.musicPlayer.Play()
}

func (gsm *GxlSoundManager) NewSound(file string) *GxlSound {
	sound := gsm.cache.LoadSound(file, cache.CacheOptions{})

	player := gsm.context.NewPlayerFromBytes(sound.data)

	return &GxlSound{player: player}
}
