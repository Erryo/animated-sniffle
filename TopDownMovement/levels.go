package main

import (
	"github.com/veandco/go-sdl2/mix"
)

func initLevel(name string, music *mix.Music) level {
	var level level
	level.name = name
	level.music = music
	level.player = nil
	level.nextID = 0
	return level
}
