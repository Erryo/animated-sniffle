package main

import (
	"fmt"
	"runtime"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	runtime.LockOSThread()
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	if err = img.Init(img.INIT_PNG); err != nil {
		panic(err)
	}
	if err = mix.Init(mix.INIT_OGG); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Sniffle Shoots Asteroids", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	if err = mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, 2, 1024); err != nil {
		panic(fmt.Errorf("error opeining audio:%v", err))
	}

	state := gameState{window: window, nextID: 1, renderer: renderer, Enemies: &[]Enemy{}, Projectiles: &[]Projectile{}, TextManager: &TextManager{}}
	state.TextManager.setDict()

	state.initPlayer(WINDOW_WIDTH/2, WINDOW_HEIGHT/2, 7, "media/player.png")
	state.loadMedia()

	state.TextManager.elements = &[]dataElement{}

	// Prepare Main Menu
	if err := state.music.Play(-1); err != nil {
		fmt.Println(err)
	}

	state.TextManager.addElement(nil, "", "Sniffle", (WINDOW_WIDTH-7*FONT_W*3)/2, 120, 3, 255, 255, 255, state.AssignID())
	state.TextManager.addElement(nil, "", "Shoots", (WINDOW_WIDTH-6*FONT_W*5)/2+2, 202, 5, 123, 123, 123, state.AssignID())
	state.TextManager.addElement(nil, "", "Shoots", (WINDOW_WIDTH-6*FONT_W*5)/2, 200, 5, 255, 255, 255, state.AssignID())
	state.TextManager.addElement(nil, "", "Asteroids", (WINDOW_WIDTH-9*FONT_W*4)/2, 320, 4, 255, 255, 255, state.AssignID())
	state.TextManager.addElement(nil, "startbutton", "Press Fire to Start", (WINDOW_WIDTH-17*FONT_W*1)/2, 460, 1, 255, 255, 255, state.AssignID())

	if state.mainMenuLoop() {
		return
	}

	state.TextManager.clearElements()
	state.gameLoop()
}
