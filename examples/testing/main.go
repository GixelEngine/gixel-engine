package main

import (
	"github.com/GixelEngine/gixel-engine/examples/testing/states"
	"github.com/GixelEngine/gixel-engine/gixel"
)

func main() {
	gixel.NewGame(640, 480, "Hello Gixel", &states.MenuState{}, 2)
}
