package main

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func initLevel(name, pathToMap string, music *mix.Music, keyDown, keyUp func(*sdl.KeyboardEvent) bool) level {
	var level level
	level.name = name
	level.music = music
	level.doKeyDown = keyDown
	level.doKeyUp = keyUp
	level.pathToMap = pathToMap
	if pathToMap != "" {
		level.loadMap()
	}
	level.player = nil
	level.nextID = 0
	return level
}

func doKeyDownWithPlayerMovement(e *sdl.KeyboardEvent) bool {
	return false
}

func doKeyUpithPlayerMovement(e *sdl.KeyboardEvent) bool {
	return false
}
