package main

import (
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
)

func main() {
	var err error
	state := initSDL()

	state.textureAtlas, err = img.LoadTexture(state.renderer, PathToTextureAtlas)
	if err != nil {
		log.Panicln("error loading textureAtlas:", err)
	}

	state.fontAtlas, err = img.LoadTexture(state.renderer, PathToFontAtlas)
	if err != nil {
		log.Panicln("error loading font Atlas:", err)
	}
	music, err := mix.LoadMUS(PathToMainMenuMusic)
	if err != nil {
		log.Panicln("error loading main Menu Music", err)
	}

	mainMenu := initLevel("MainMenu", music)
	state.currentLevel = &mainMenu
	state.levels = append(state.levels, &mainMenu)

	for i := range MaxLevelHeigth - 1 {
		state.currentLevel.tiles[i][i] = TilesStartIndex
	}
	state.currentLevel.tiles[MaxViewHeigth-1][MaxViewWidth-1] = TilesStartIndex + 1
	state.currentLevel.tiles[MaxViewHeigth][MaxViewWidth] = TilesStartIndex + 3
	//	state.currentLevel.tiles[MaxLevelHeigth-1][MaxLevelWidth-1] = TilesStartIndex + 2
	state.gameLoop()
}
