package main

import (
	"fmt"
	"strconv"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func (state *state) drawEnemies() {
	for _, enemy := range *state.currentLevel.enemies {
		state.renderer.SetDrawColor(enemy.color[0], enemy.color[1], enemy.color[2], 255)
		state.renderer.FillRect(enemy.rect)
		gfx.CircleRGBA(state.renderer, enemy.x, enemy.y, int32(enemy.hitBoxRadius), 255, 0, 0, 255)
	}
}

func (state *state) drawProjectiles() {
	for _, projectile := range *state.currentLevel.projectiles {
		state.renderer.SetDrawColor(projectile.color[0], projectile.color[1], projectile.color[2], 255)
		state.renderer.FillRect(projectile.rect)
		gfx.CircleRGBA(state.renderer, projectile.x, projectile.y, int32(projectile.hitBoxRadius), 255, 0, 0, 255)
	}
}

func (state *state) drawUI() {
	state.drawElements()

	ammoText := "Ammo:" + strconv.Itoa(int(state.currentLevel.player.ammo))
	rect := sdl.Rect{X: 0, Y: 0, W: int32(len(ammoText) * FONT_W * 2), H: FONT_W * 2}
	state.renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	state.renderer.SetDrawColor(123, 123, 123, 120)
	state.renderer.FillRect(&rect)
	// state.TextManager.print(state.renderer, ammoText, 2, 0, 0, 255, 255, 255)
}

func (lvl *level) drawAllObjects(state *state) {
	state.drawEnemies()
	state.drawProjectiles()
	state.drawUI()
	state.renderer.Present()
}

func (state *state) prepareScene(backgroundImage *sdl.Texture) {
	state.renderer.SetDrawColor(255, 255, 255, 255)
	state.renderer.Clear()
	if err := state.renderer.Copy(backgroundImage, nil, nil); err != nil {
		fmt.Println("Error copying backgroundImage: ", err)
	}
}

func (state *state) blit(player player) {
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
