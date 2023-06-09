package main

import (
	"embed"

	"github.com/GixelEngine/gixel-engine/examples/testing/states"
	"github.com/GixelEngine/gixel-engine/gixel"
)

//go:embed assets
var assets embed.FS

func main() {
	gixel.NewGame(640, 480, "Hello Gixel", &assets, &states.MenuState{}, 2).Run()
}
