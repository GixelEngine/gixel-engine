package main

import (
	"embed"

	"github.com/GixelEngine/gixel-engine/gixel"
)

/*
	ported from: https://github.com/sedyh/ebitengine-bunny-mark
*/

const GAME_WIDTH = 800
const GAME_HEIGHT = 600

//go:embed assets
var assets embed.FS

func main() {
	gixel.NewGame(GAME_WIDTH, GAME_HEIGHT, "Gixel Bunnymark", &assets, &PlayState{}, 1).Run()
}
