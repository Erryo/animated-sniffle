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

		state.loadMedia()
		state.prepareScene()
		state.Update()
		state.drawAllGameObjects()
		sdl.Delay(32)
	}
}

func (state *gameState) loadMedia() {
	var err error
	if state.backgroundImage, err = img.LoadTexture(state.renderer, "media/background.png"); err != nil {
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
	speed := state.Player.speed
	switch event.Keysym.Scancode {
	case sdl.SCANCODE_W:
		if state.Player.y-int(speed) > 0+20 {
			state.Player.y += -int(speed)
		}
	case sdl.SCANCODE_A:
		if state.Player.x-int(speed) > 0+20 {
			state.Player.x += -int(speed)
		}
	case sdl.SCANCODE_S:
		if state.Player.y+int(speed) < WINDOW_HEIGHT-80 {
			state.Player.y += int(speed)
		}
	case sdl.SCANCODE_D:
		if state.Player.x+int(speed) < WINDOW_WIDTH-80 {
			state.Player.x += int(speed)
		}
	}
}

func (state *gameState) doKeyUp(event *sdl.KeyboardEvent) {
}

func (state *gameState) QuitGame() {
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

func (state *gameState) AssignID() uint16 {
	state.nextID++
	return state.nextID - 1
}

func (state *gameState) initObject(color [3]uint8, x, y, w, h int32) {
	rect := sdl.Rect{X: x, Y: y, W: w, H: h}
	object := Enemy{rect: rect, color: color, id: state.AssignID()}
	*state.Enemies = append(*state.Enemies, object)
}

// texturePath := ,,media/name.png
func (state *gameState) initPlayer(x, y int, texturePath string) {
	texture, err := img.LoadTexture(state.renderer, texturePath)
	if err != nil {
		panic(err)
	}
	player := Player{x: x, y: y, texture: texture, id: state.AssignID(), speed: 20}
	state.Player = &player
}

func (state *gameState) Update() {
	state.blit(*state.Player)
}
