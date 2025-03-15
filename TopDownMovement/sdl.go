package main

import (
	"log"
	"runtime"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func initSDL() state {
	runtime.LockOSThread()
	var err error
	var state state
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Panicln(err)
	}
	defer sdl.Quit()

	if err = img.Init(img.INIT_PNG); err != nil {
		log.Panicln(err)
	}
	if err = mix.Init(mix.INIT_OGG); err != nil {
		log.Panicln(err)
	}

	window, err := sdl.CreateWindow("Top Down Sniffle", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		WindowWidth, WindowHeigth, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Panicln(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Panicln(err)
	}

	if err = mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, 2, 1024); err != nil {
		log.Panicln("error opening music: ", err)
	}

	state.window = window
	state.renderer = renderer
	return state
}

func (s *state) quitGame() {
	for i := range s.levels {
		s.levels[i].destroy()
		s.levels[i] = nil
		s.currentLevel = nil
	}
	s.fontAtlas.Destroy()
	s.renderer.Destroy()
	s.renderer = nil
	s.window.Destroy()
	s.window = nil
}

func (l *level) destroy() {
	l.backgroundImage.Destroy()
	var zLayer []bliter
	for i := range l.blitables {
		zLayer = l.blitables[i]
		for j := range zLayer {
			zLayer[j].destroy()
		}
	}
	l = &(level{})
}

func (s *sprite) blit() {
}

func (s *sprite) destroy() {
}

func closeSDL() {
	mix.CloseAudio()
	mix.Quit()
	img.Quit()
	sdl.Quit()
}
