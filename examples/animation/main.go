package main

import (
	"github.com/odedro987/gixel-engine/gixel"
)

func main() {
	gixel.NewGame(640, 480, "Gixel Animation", &PlayState{}, 2)
}
