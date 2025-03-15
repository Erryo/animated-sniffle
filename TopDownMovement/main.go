package main

import (
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
)

func main() {
	var err error
	state := initSDL()

	music, err := mix.LoadMUS(PathToMainMenuMusic)
	if err != nil {
		log.Panicln("error loading main Menu Music", err)
	}
	mainMenu := initLevel("MainMenu", music)
	state.currentLevel = &mainMenu
	state.levels = append(state.levels, &mainMenu)

	state.textureAtlas, err = img.LoadTexture(state.renderer, PathToTextureAtlas)
	if err != nil {
		log.Panicln("error loading textureAtlas:", err)
	}

	state.fontAtlas, err = img.LoadTexture(state.renderer, PathToFontAtlas)
	if err != nil {
		log.Panicln("error loading font Atlas:", err)
	}
}
