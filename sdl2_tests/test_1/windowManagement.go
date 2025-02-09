package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func (state *gameState) drawAllGameObjects() {
	for _, enemy := range *state.Enemies {
		state.renderer.SetDrawColor(enemy.color[0], enemy.color[1], enemy.color[2], 255)
		state.renderer.FillRect(&enemy.rect)
		gfx.CircleRGBA(state.renderer, enemy.x, enemy.y, int32(enemy.hitBoxRadius), 255, 0, 0, 255)
		state.renderer.DrawLine(enemy.x, enemy.y, state.Player.x, state.Player.y)
	}
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
	// The origin is the same but the texture is centered
	dst.X = dst.X - dst.W/2
	dst.Y = dst.Y - dst.H/2
	state.renderer.Copy(player.texture, nil, &dst)

	gfx.CircleRGBA(state.renderer, player.x, player.y, int32(player.hitBoxRadius), 255, 0, 0, 255)
}
