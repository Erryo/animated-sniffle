package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func (state *gameState) drawEnemies() {
	for _, enemy := range *state.Enemies {
		state.renderer.SetDrawColor(enemy.color[0], enemy.color[1], enemy.color[2], 255)
		state.renderer.FillRect(enemy.rect)
		gfx.CircleRGBA(state.renderer, enemy.x, enemy.y, int32(enemy.hitBoxRadius), 255, 0, 0, 255)
	}
}

func (state *gameState) drawProjectiles() {
	for _, projectile := range *state.Projectiles {
		state.renderer.SetDrawColor(projectile.color[0], projectile.color[1], projectile.color[2], 255)
		state.renderer.FillRect(projectile.rect)
		gfx.CircleRGBA(state.renderer, projectile.x, projectile.y, int32(projectile.hitBoxRadius), 255, 0, 0, 255)
	}
}

func (state *gameState) drawAllObjects() {
	state.drawEnemies()
	state.drawProjectiles()
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
	if err := state.renderer.CopyEx(player.texture, nil, &dst, float64(player.rotation), nil, sdl.FLIP_NONE); err != nil {
		fmt.Println("Error CopyEx'ng player:", err)
	}

	gfx.CircleRGBA(state.renderer, player.x, player.y, int32(player.hitBoxRadius), 255, 0, 0, 255)
}
