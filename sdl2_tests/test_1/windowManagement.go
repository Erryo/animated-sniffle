package main

import "fmt"

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
