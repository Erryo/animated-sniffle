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
	// state.TextManager.addElement(&state.currentLevel.player.ammo, "", "Ammo:", 0, 0, 2, 255, 255, 255, state.AssignID())
	elem, err := state.currentLevel.getElementByName("frameTime")
	if err != nil {
		fmt.Println(err)
	}
	elem.data = &frameTime

outerGameLoop:
	for running {

		start = time.Now()
		if state.doInput() {
			running = false
			break outerGameLoop
		}

		state.prepareScene(state.currentLevel.backgroundImage)
		state.currentLevel.Update(state)
		frameTime = time.Since(start).String()
		state.currentLevel.drawAllObjects(state)

		sdl.Delay(GAME_UPDATE_DELAY)
	}
}

func (state *state) mainMenuLoop() bool {
	running := true

	var startTextY int32
	var offset int32
	offset = 1
	startTextY = 460
	animatedStartText, err := state.currentLevel.getElementByName("startbutton")
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

		state.prepareScene(state.currentLevel.backgroundImage)
		if animatedStartText != nil {
			if startTextY > 470 || startTextY < 450 {
				offset = offset * -1
			}
			startTextY += offset
			animatedStartText.y = startTextY

		}
		state.drawElements()
		state.renderer.Present()
		sdl.Delay(GAME_UPDATE_DELAY)
	}
	return true
}

func (state *state) loadMedia() {
	var err error

	if state.fontMap, err = img.LoadTexture(state.renderer, "media/fontmap.png"); err != nil {
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
		state.currentLevel.player.eventList[0] = true

	case sdl.SCANCODE_A:
		state.currentLevel.player.eventList[1] = true

	case sdl.SCANCODE_S:
		state.currentLevel.player.eventList[2] = true

	case sdl.SCANCODE_D:
		state.currentLevel.player.eventList[3] = true

	case sdl.SCANCODE_SPACE:
		state.currentLevel.player.eventList[4] = true
	case sdl.SCANCODE_Q:
		state.currentLevel.player.eventList[5] = true
	case sdl.SCANCODE_E:
		state.currentLevel.player.eventList[6] = true
	case sdl.SCANCODE_ESCAPE:
		return state.switchToMainMenu()
	}
	return false
}

func (state *state) doKeyUp(event *sdl.KeyboardEvent) {
	switch event.Keysym.Scancode {
	case sdl.SCANCODE_W:
		state.currentLevel.player.eventList[0] = false

	case sdl.SCANCODE_A:
		state.currentLevel.player.eventList[1] = false
	case sdl.SCANCODE_S:
		state.currentLevel.player.eventList[2] = false

	case sdl.SCANCODE_D:
		state.currentLevel.player.eventList[3] = false
	case sdl.SCANCODE_SPACE:
		state.currentLevel.player.eventList[4] = false
	case sdl.SCANCODE_Q:
		state.currentLevel.player.eventList[5] = false
	case sdl.SCANCODE_E:
		state.currentLevel.player.eventList[6] = false

	}
}

func (state *state) switchToMainMenu() bool {
	mainMenu := state.getLevelByName("mainMenu")
	if mainMenu == nil {
		return false
	}

	state.currentLevel = mainMenu

	if state.mainMenuLoop() {
		fmt.Println("main menu quit")
		return true
	}
	state.switchToGameLevel()
	return false
}

func (state *state) switchToGameLevel() {
	gameLvl := state.getLevelByName("game")
	gameLvl.music = state.currentLevel.music
	state.currentLevel = gameLvl
}

func (state *state) createGameLevel() {
	var gameLvl level
	var err error

	gameLvl.name = "game"
	gameLvl.enemies = &[]enemy{}
	gameLvl.projectiles = &[]projectile{}
	gameLvl.dataElements = &[]dataElement{}
	gameLvl.nextID = 1

	gameLvl.initPlayer(WINDOW_WIDTH/2, WINDOW_HEIGHT/2, 7, "media/player.png", state.renderer)

	if err = gameLvl.addElement(nil, "frameTime", "fT:", WINDOW_WIDTH-300, 0, 1, WHITE, state.currentLevel.AssignID()); err != nil {
		fmt.Println(err)
	}
	if err = gameLvl.addElement(&gameLvl.player.ammo, "", "Ammo:", 0, 0, 2, WHITE, state.currentLevel.AssignID()); err != nil {
		fmt.Println(err)
	}

	if gameLvl.backgroundImage, err = img.LoadTexture(state.renderer, "media/background.png"); err != nil {
		panic(err)
	}
	if gameLvl.music, err = mix.LoadMUS("media/music.ogg"); err != nil {
		panic(err)
	}

	state.levels = append(state.levels, &gameLvl)
}

func (state *state) createMainMenu() {
	var mainMenu level
	var err error

	mainMenu.name = "mainMenu"
	mainMenu.enemies = &[]enemy{}
	mainMenu.projectiles = &[]projectile{}
	mainMenu.dataElements = &[]dataElement{}
	mainMenu.nextID = 1

	mainMenu.addElement(nil, "", "Sniffle", (WINDOW_WIDTH-7*FONT_W*3)/2, 120, 3, WHITE, mainMenu.AssignID())
	mainMenu.addElement(nil, "", "Shoots", (WINDOW_WIDTH-6*FONT_W*5)/2+2, 202, 5, [3]uint8{123, 123, 123}, mainMenu.AssignID())
	mainMenu.addElement(nil, "", "Shoots", (WINDOW_WIDTH-6*FONT_W*5)/2, 200, 5, WHITE, mainMenu.AssignID())
	mainMenu.addElement(nil, "", "Asteroids", (WINDOW_WIDTH-9*FONT_W*4)/2, 320, 4, WHITE, mainMenu.AssignID())
	mainMenu.addElement(nil, "startbutton", "Press Fire to Start", (WINDOW_WIDTH-17*FONT_W*1)/2, 460, 1, WHITE, mainMenu.AssignID())

	if mainMenu.backgroundImage, err = img.LoadTexture(state.renderer, "media/background.png"); err != nil {
		panic(err)
	}
	if mainMenu.music, err = mix.LoadMUS("media/music.ogg"); err != nil {
		panic(err)
	}

	state.levels = append(state.levels, &mainMenu)
}

func (state *state) QuitGame() {
	mix.HaltChannel(-1)
	mix.HaltMusic()
	state.currentLevel.music.Free()
	state.currentLevel.music = nil
	state.currentLevel.player.shootEff.Free()
	state.currentLevel.player.shootEff = nil
	state.currentLevel.player.texture.Destroy()
	state.currentLevel.player.texture = nil
	state.fontMap.Destroy()
	state.fontMap = nil
	state.currentLevel.backgroundImage.Destroy()
	state.currentLevel.backgroundImage = nil
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

func (state *state) getLevelByName(name string) *level {
	for idx, lvl := range state.levels {
		if lvl.name == name {
			return state.levels[idx]
		}
	}
	return nil
}

func (level *level) Update(state *state) {
	level.spawnEnemies()
	level.player.checkEventList(level)
	// needs Testing <- written in a hurry

	level.player.handleFireCooldown()
	level.moveEnemies()
	level.moveProjectiles()
	state.blit(*level.player)
}
