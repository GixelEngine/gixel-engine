package main

import (
	"embed"

	"github.com/GixelEngine/gixel-engine/gixel"
)

//go:embed assets
var assets embed.FS

func main() {
	gixel.NewGame(640, 480, "Gixel Animation", &assets, &PlayState{}, 2).Run()
}
