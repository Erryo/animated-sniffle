package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	if err := img.Init(img.INIT_PNG); err != nil {
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

	state := gameState{window: window, nextID: 1, renderer: renderer, Enemies: &[]Enemy{}, Projectiles: &[]Projectile{}}

	state.initPlayer(WINDOW_WIDTH/2, WINDOW_HEIGHT/2, 10, "media/player.png")
	state.initObject(CYAN, 600, 200, 40, 40, 5)

	state.loadMedia()

	state.gameLoop()
}
