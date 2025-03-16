package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func (s *state) gameLoop() {
	var previousFrame uint64
	var currentFrame uint64

	running := true
	for running {
		previousFrame = currentFrame
		currentFrame = sdl.GetPerformanceCounter()
		s.deltaTime = float32((currentFrame-previousFrame)*1000) / float32(sdl.GetPerformanceFrequency())

		if s.handleInput() {
			s.quitGame()
			closeSDL()
			return
		}

		s.prepareScene()
		s.drawTiles()

		s.renderer.Present()
	}
}

// returns true if game should quit
func (s *state) handleInput() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			return true
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				return s.currentLevel.doKeyDown(e)
			} else if e.Type == sdl.KEYUP {
				return s.currentLevel.doKeyDown(e)
			}

		}
	}
	return false
}
