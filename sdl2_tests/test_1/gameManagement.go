package main

import (
	"fmt"
	"math"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (state *gameState) gameLoop() {
	running := true
	var Start time.Time
	gameStart := time.Now()
outerGameLoop:
	for running {
		Start = time.Now()
		if state.doInput() {
			running = false
			break outerGameLoop
		}

		state.prepareScene()
		state.Update()
		state.drawAllGameObjects()

		fmt.Println(time.Since(Start))
		if time.Since(Start) > 50*time.Millisecond {
			fmt.Println(time.Since(gameStart))
		}
		sdl.Delay(22)
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

func (state *gameState) AssignID() uint16 {
	state.nextID++
	return state.nextID - 1
}

func (state *gameState) initObject(color [3]uint8, x, y, w, h int32) {
	rect := sdl.Rect{X: x - w/2, Y: y - h/2, W: w, H: h}
	radius := math.Sqrt(math.Pow(float64(h/2), 2) + math.Pow(float64(w/2), 2))
	object := Enemy{x: x, y: y, rect: rect, color: color, id: state.AssignID(), hitBoxRadius: uint8(radius)}
	*state.Enemies = append(*state.Enemies, object)
}

// texturePath := ,,media/name.png
func (state *gameState) initPlayer(x, y int32, speed uint8, texturePath string) {
	texture, err := img.LoadTexture(state.renderer, texturePath)
	if err != nil {
		panic(err)
	}
	eventList := make([]bool, 6)

	_, _, w, h, err := texture.Query()
	if err != nil {
		panic(err)
	}

	radius := math.Sqrt(math.Pow(float64(h/2), 2) + math.Pow(float64(w/2), 2))
	player := Player{x: x, y: y, texture: texture, id: state.AssignID(), speed: speed, eventList: eventList, hitBoxRadius: uint8(radius)}
	state.Player = &player
}

func (state *gameState) Update() {
	state.Player.checkEventList(state)
	state.blit(*state.Player)
}
