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
	context       *audio.Context
	musicPlayer   *audio.Player
	isMusicPaused bool
	isMusicLooped bool
	cache         *GxlSoundCache
}

func NewSoundManager(fs *embed.FS) *GxlSoundManager {
	return &GxlSoundManager{
		context: audio.NewContext(sampleRate),
		cache:   NewSoundCache(fs),
	}
}

func (gsm *GxlSoundManager) PlayMusic(file string, volume float64, loop bool) {
	sound := gsm.cache.LoadSound(file, cache.CacheOptions{})

	gsm.musicPlayer = gsm.context.NewPlayerFromBytes(sound.data)
	gsm.isMusicPaused = false
	gsm.isMusicLooped = loop
	gsm.musicPlayer.SetVolume(volume)
	gsm.musicPlayer.Play()
}

func (gsm *GxlSoundManager) ResumeMusic() {
	if gsm.musicPlayer == nil || gsm.musicPlayer.IsPlaying() {
		return
	}

	gsm.isMusicPaused = false
	gsm.musicPlayer.Play()
}

func (gsm *GxlSoundManager) PauseMusic() {
	gsm.isMusicPaused = true
	gsm.musicPlayer.Pause()
}

func (gsm *GxlSoundManager) SetMusicVolume(volume float64) {
	if gsm.musicPlayer == nil {
		return
	}

	gsm.musicPlayer.SetVolume(volume)
}

func (gsm *GxlSoundManager) IsMusicPaused() bool {
	return gsm.isMusicPaused
}

func (gsm *GxlSoundManager) NewSound(file string) *GxlSound {
	sound := gsm.cache.LoadSound(file, cache.CacheOptions{})

	player := gsm.context.NewPlayerFromBytes(sound.data)

	return &GxlSound{player: player}
}

func (gsm *GxlSoundManager) Update() {
	if gsm.musicPlayer == nil {
		return
	}

	if gsm.isMusicLooped && !gsm.musicPlayer.IsPlaying() && !gsm.isMusicPaused {
		gsm.isMusicPaused = false
		gsm.musicPlayer.Rewind()
		gsm.musicPlayer.Play()
	}
}
