package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (state *gameState) gameLoop() {
	running := true
outerGameLoop:
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				state.QuitGame()
				state.CloseSDL()
				running = false
				break outerGameLoop
			}
		}
		state.loadMedia()
		state.prepareScene()
		state.Update()
		state.drawAllGameObjects()
		sdl.Delay(16)
	}
}

func (state *gameState) loadMedia() {
	var err error
	if state.backgroundImage, err = img.LoadTexture(state.renderer, "media/background.png"); err != nil {
		panic(err)
	}
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

func (state *gameState) initObject(pixel uint32, x, y, w, h int32) {
	rect := sdl.Rect{X: x, Y: y, W: w, H: h}
	object := Object{rect: rect, pixel: pixel, id: state.nextID}
	*state.gameObjects = append(*state.gameObjects, object)
	state.nextID += 1
}

func (state *gameState) Update() {
	state.renderer.SetDrawColor(0, 0, 0, 255)
}
