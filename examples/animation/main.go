package main

import (
	"github.com/GixelEngine/gixel-engine/gixel"
)

func main() {
	gixel.NewGame(640, 480, "Gixel Animation", &PlayState{}, 2)
}
