package main

import (
	"log"
	"os"

	leveleditor "github.com/Erryo/TopDownSniffle/levelEditor"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
)

func main() {
	var err error
	state := initSDL()
	args := os.Args[1:]
	if len(args) > 0 {
		if state.handleArgs(args) {
			closeSDL()
			return
		}
	}

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

	mainMenu := initLevel("MainMenu", PathToFirstMap, music, doKeyDownWithPlayerMovement, doKeyUpithPlayerMovement)
	state.currentLevel = &mainMenu
	state.levels = append(state.levels, &mainMenu)

	for i := range MaxLevelHeigth - 1 {
		state.currentLevel.tiles[i][0] = TilesStartIndex + uint8(i%5)
	}
	state.currentLevel.tiles[MaxViewHeigth-1][MaxViewWidth-1] = TilesStartIndex + 1
	state.currentLevel.tiles[MaxViewHeigth][MaxViewWidth] = TilesStartIndex + 3
	//	state.currentLevel.tiles[MaxLevelHeigth-1][MaxLevelWidth-1] = TilesStartIndex + 2
	state.gameLoop()
}

func (s *state) handleArgs(args []string) bool {
	switch args[0] {
	case "edit":
		s.window.Hide()
		leveleditor.StartLevelEditor()
		return true
	}
	return false
}
