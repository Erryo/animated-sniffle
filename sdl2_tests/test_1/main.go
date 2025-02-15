package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
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

	if err := state.music.Play(-1); err != nil {
		fmt.Println(err)
	}
	if state.mainMenuLoop() {
		return
	}
	state.gameLoop()
}
