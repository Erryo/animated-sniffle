package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func (state *gameState) drawAllGameObjects() {
	state.renderer.Present()
}

func (state *gameState) prepareScene() {
	state.renderer.SetDrawColor(255, 255, 255, 255)
	state.renderer.Clear()
	if err := state.renderer.Copy(state.backgroundImage, nil, nil); err != nil {
		fmt.Println("Error copying backgroundImage: ", err)
	}
}

func (state *gameState) blit(player Player) {
	var err error
	dst := sdl.Rect{X: int32(player.x), Y: int32(player.y)}
	_, _, dst.W, dst.H, err = player.texture.Query()
	if err != nil {
		fmt.Println("Error blint´ng player:", err)
	}
	state.renderer.Copy(player.texture, nil, &dst)
}
