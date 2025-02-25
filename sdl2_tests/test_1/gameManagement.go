package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

// Rework is needed
// Level struct needed (just because of TextELements YAYYY)
// basically separate the state from each level(menu)
// OVERKILL
// THIS GAME ONLY NEEDS 2 MENUS

func (state *state) gameLoop() {
	running := true

	var start time.Time
	var frameTime string

	// state.TextManager.addElement(&frameTime, "", "fT:", WINDOW_WIDTH-300, 0, 1, 255, 255, 255, state.AssignID())
	// state.TextManager.addElement(&state.Player.ammo, "", "Ammo:", 0, 0, 2, 255, 255, 255, state.AssignID())

outerGameLoop:
	for running {

		start = time.Now()
		if state.doInput() {
			running = false
			break outerGameLoop
		}

		state.prepareScene()
		state.Update()
		frameTime = time.Since(start).String()
		//		state.TextManager.print(state.renderer, "fT:"+time.Since(start).String(), 1, WINDOW_WIDTH-300, 0, 255, 255, 255)
		state.drawAllObjects()

		sdl.Delay(GAME_UPDATE_DELAY)
	}
}

func (state *state) mainMenuLoop() bool {
	running := true
	var startTextY int32
	var offset int32
	offset = 1
	startTextY = 460
	// animatedStartText, err := state.TextManager.getElementByName("startbutton")
	if err != nil {
		fmt.Println(err)
	}

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
		if animatedStartText != nil {
			if startTextY > 470 || startTextY < 450 {
				offset = offset * -1
			}
			startTextY += offset
			animatedStartText.y = startTextY

		}
		state.TextManager.drawElements(state.renderer)
		state.renderer.Present()
		sdl.Delay(GAME_UPDATE_DELAY)
	}
	return true
}

func (state *state) loadMedia() {
	var err error
	if state.backgroundImage, err = img.LoadTexture(state.renderer, "media/background.png"); err != nil {
		panic(err)
	}
	if state.TextManager.fontMap, err = img.LoadTexture(state.renderer, "media/fontmap.png"); err != nil {
		panic(err)
	}
	if state.music, err = mix.LoadMUS("media/music.ogg"); err != nil {
		panic(err)
	}
	if state.Player.shootEff, err = mix.LoadWAV("media/shoot.ogg"); err != nil {
		panic(err)
	}
}

// returns true if recieved quit signal
func (state *state) doInput() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			state.QuitGame()
			state.CloseSDL()
			return true
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				return state.doKeyDown(e)
			} else if e.Type == sdl.KEYUP {
				state.doKeyUp(e)
			}

		}
	}
	return false
}

func (state *state) doKeyDown(event *sdl.KeyboardEvent) bool {
	if event.Repeat != 0 {
		return false
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
	case sdl.SCANCODE_ESCAPE:
		return state.switchToMainMenu()
	}
	return false
}

func (state *state) doKeyUp(event *sdl.KeyboardEvent) {
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

func (state *state) switchToMainMenu() bool {
	state.TextManager.clearElements()
	state.TextManager.addElement(nil, "", "Sniffle", (WINDOW_WIDTH-7*FONT_W*3)/2, 120, 3, 255, 255, 255, state.AssignID())
	state.TextManager.addElement(nil, "", "Shoots", (WINDOW_WIDTH-6*FONT_W*5)/2+2, 202, 5, 123, 123, 123, state.AssignID())
	state.TextManager.addElement(nil, "", "Shoots", (WINDOW_WIDTH-6*FONT_W*5)/2, 200, 5, 255, 255, 255, state.AssignID())
	state.TextManager.addElement(nil, "", "Asteroids", (WINDOW_WIDTH-9*FONT_W*4)/2, 320, 4, 255, 255, 255, state.AssignID())
	state.TextManager.addElement(nil, "startbutton", "Press Fire to Start", (WINDOW_WIDTH-17*FONT_W*1)/2, 460, 1, 255, 255, 255, state.AssignID())

	if state.mainMenuLoop() {
		return true
	}
	state.TextManager.clearElements()
	state.switchToGameLevel()
	return false
}

func (state *state) switchToGameLevel() {
	state.TextManager.addElement(&state.Player.ammo, "", "Ammo:", 0, 0, 2, 255, 255, 255, state.AssignID())
}

func (state *state) QuitGame() {
	mix.HaltChannel(-1)
	mix.HaltMusic()
	state.music.Free()
	state.music = nil
	state.Player.shootEff.Free()
	state.Player.shootEff = nil
	state.Player.texture.Destroy()
	state.Player.texture = nil
	state.TextManager.fontMap.Destroy()
	state.TextManager.fontMap = nil
	state.backgroundImage.Destroy()
	state.backgroundImage = nil
	state.renderer.Destroy()
	state.renderer = nil
	state.window.Destroy()
	state.window = nil
}

func (state *state) CloseSDL() {
	mix.CloseAudio()
	mix.Quit()
	img.Quit()
	sdl.Quit()
}

func (state *state) Update() {
	state.spawnEnemies()
	state.Player.checkEventList(state)
	// needs Testing <- written in a hurry

	state.Player.handleFireCooldown()
	state.moveEnemies()
	state.moveProjectiles()
	state.blit(*state.Player)
}
