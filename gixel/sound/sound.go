package sound

import (
	"embed"
	"log"

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

func (gs *GxlSound) Play(loop bool) {
	if gs.player.IsPlaying() {
		return
	}

	gs.player.Rewind()
	gs.player.Play()
}

func (gs *GxlSound) SetVolume(volume float64) {
	if volume < 0 || volume > 1 {
		log.Panicln("volume must be between 0 and 1")
	}
	gs.player.SetVolume(volume)
}

func (gs *GxlSound) GetVolume() float64 {
	return gs.player.Volume()
}

type GxlSoundManager struct {
	context        *audio.Context
	musicPlayer    *audio.Player
	isMusicPlaying bool
	isMusicPaused  bool
	isMusicLooped  bool
	cache          *GxlSoundCache
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
	gsm.isMusicPlaying = true
	gsm.isMusicPaused = false
	gsm.isMusicLooped = loop
	gsm.musicPlayer.SetVolume(volume)
	gsm.musicPlayer.Play()
}

func (gsm *GxlSoundManager) ResumeMusic() {
	if gsm.musicPlayer == nil || gsm.musicPlayer.IsPlaying() {
		return
	}

	gsm.isMusicPlaying = true
	gsm.isMusicPaused = false
	gsm.musicPlayer.Play()
}

func (gsm *GxlSoundManager) PauseMusic() {
	gsm.isMusicPlaying = false
	gsm.isMusicPaused = true
	gsm.musicPlayer.Pause()
}

func (gsm *GxlSoundManager) SetMusicVolume(volume float64) {
	if gsm.musicPlayer == nil {
		return
	}

	gsm.musicPlayer.SetVolume(volume)
}

func (gsm *GxlSoundManager) IsMusicPlaying() bool {
	return gsm.isMusicPlaying
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
		gsm.isMusicPlaying = true
		gsm.isMusicPaused = false
		gsm.musicPlayer.Rewind()
		gsm.musicPlayer.Play()
	}
}
