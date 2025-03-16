package leveleditor

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type state struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	tiles    [MaxLevelHeigth][MaxLevelWidth]uint8
}

const (
	WindowWidth, WindowHeigth = 900, 900
	MaxLevelHeigth            = 100
	MaxLevelWidth             = 100
	ArtSize                   = 16
	TextureAtlasW             = 6
	TextureAtlasH             = 9
)

func StartLevelEditor() {
	window, err := sdl.CreateWindow("level editor", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		WindowWidth, WindowHeigth, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Panicln(err)
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Panicln(err)
	}
	state := state{renderer: renderer, window: window}
	state.gameLoop()
}

func (s *state) gameLoop() {
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("quitting")
				running = false
				return
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
				} else if e.Type == sdl.KEYUP {
				}

			}
		}
		s.renderer.SetDrawColor(255, 255, 255, 255)
		s.renderer.Clear()
		s.renderer.Present()
	}
}

func indexToPosition(idx uint8) (uint8, uint8) {
	x := (idx) % TextureAtlasW
	y := (idx) / TextureAtlasW
	return x, y
}
