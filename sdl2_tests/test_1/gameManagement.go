package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (state *gameState) gameLoop() {
	running := true

outerGameLoop:
	for running {

		if state.doInput() {
			running = false
			break outerGameLoop
		}

		state.prepareScene()
		state.Update()
		state.drawAllObjects()

		sdl.Delay(GAME_UPDATE_DELAY)
	}
}

func (state *gameState) mainMenuLoop() bool {
	running := true
	var startTextY int32
	var offset int32
	offset = 1
	startTextY = 460
	for running {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				state.QuitGame()
				state.CloseSDL()
				return true
				running = false
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					switch e.Keysym.Scancode {
					case sdl.SCANCODE_SPACE:
						return false
					}
				}

			}
		}

		state.prepareScene()
		state.TextManager.print(state.renderer, "Sniffle", 3, (WINDOW_WIDTH-7*FONT_W*3)/2, 120, 255, 255, 255)

		state.TextManager.print(state.renderer, "Shoots", 5, (WINDOW_WIDTH-6*FONT_W*5)/2+2, 202, 123, 123, 123)
		state.TextManager.print(state.renderer, "Shoots", 5, (WINDOW_WIDTH-6*FONT_W*5)/2, 200, 255, 255, 255)

		state.TextManager.print(state.renderer, "Asteroids", 4, (WINDOW_WIDTH-9*FONT_W*4)/2, 320, 255, 255, 255)
		if startTextY > 470 || startTextY < 450 {
			offset = offset * -1
		}
		startTextY += offset

		state.TextManager.print(state.renderer, "Press Fire to Start", 1, (WINDOW_WIDTH-17*FONT_W*1)/2, startTextY, 255, 255, 255)
		state.renderer.Present()
		sdl.Delay(GAME_UPDATE_DELAY)
	}
	return true
}

func (state *gameState) loadMedia() {
	var err error
	if state.backgroundImage, err = img.LoadTexture(state.renderer, "media/background.png"); err != nil {
		panic(err)
	}
	if state.TextManager.fontMap, err = img.LoadTexture(state.renderer, "media/fontmap.png"); err != nil {
		panic(err)
	}
}

// returns true if recieved quit signal
func (state *gameState) doInput() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			state.QuitGame()
			state.CloseSDL()
			return true
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				state.doKeyDown(e)
			} else if e.Type == sdl.KEYUP {
				state.doKeyUp(e)
			}

		}
	}
	return false
}

func (state *gameState) doKeyDown(event *sdl.KeyboardEvent) {
	if event.Repeat != 0 {
		return
	}
	switch event.Keysym.Scancode {
	case sdl.SCANCODE_W:
		state.Player.eventList[0] = true

	case sdl.SCANCODE_A:
		state.Player.eventList[1] = true

	case sdl.SCANCODE_S:
		state.Player.eventList[2] = true

	case sdl.SCANCODE_D:
		state.Player.eventList[3] = true

	case sdl.SCANCODE_SPACE:
		state.Player.eventList[4] = true
	case sdl.SCANCODE_Q:
		state.Player.eventList[5] = true
	case sdl.SCANCODE_E:
		state.Player.eventList[6] = true
	}
}

func (state *gameState) doKeyUp(event *sdl.KeyboardEvent) {
	switch event.Keysym.Scancode {
	case sdl.SCANCODE_W:
		state.Player.eventList[0] = false

	case sdl.SCANCODE_A:
		state.Player.eventList[1] = false
	case sdl.SCANCODE_S:
		state.Player.eventList[2] = false

	case sdl.SCANCODE_D:
		state.Player.eventList[3] = false
	case sdl.SCANCODE_SPACE:
		state.Player.eventList[4] = false
	case sdl.SCANCODE_Q:
		state.Player.eventList[5] = false
	case sdl.SCANCODE_E:
		state.Player.eventList[6] = false

	}
}

func (state *gameState) QuitGame() {
	state.Player.texture.Destroy()
	state.Player.texture = nil
	state.backgroundImage.Destroy()
	state.backgroundImage = nil
	state.renderer.Destroy()
	state.renderer = nil
	state.window.Destroy()
	state.window = nil
}

func (state *gameState) CloseSDL() {
	img.Quit()
	sdl.Quit()
}

func (state *gameState) Update() {
	state.spawnEnemies()
	state.Player.checkEventList(state)
	// needs Testing <- written in a hurry

	state.Player.handleFireCooldown()
	state.moveEnemies()
	state.moveProjectiles()
	state.blit(*state.Player)
}
