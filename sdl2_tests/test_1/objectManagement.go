package main

import (
	"fmt"
	"math"
)

func (state *gameState) moveProjectiles() {
	for idx, projectile := range *state.Projectiles {
		newX := projectile.x + int32(projectile.scaler[0])
		newY := projectile.y + int32(projectile.scaler[1])
		willColide, enemy := willCollide(state, newX, newY, projectile.hitBoxRadius)
		if willColide {
			if enemy != nil {
				if enemy.hp <= projectile.damage {
					fmt.Println(enemy)
					state.destroyEnemy(enemy)
				} else {
					(*enemy).hp -= 1
					fmt.Println(enemy.hp)

				}
			}
		}
		if newX+int32(projectile.hitBoxRadius) > WINDOW_WIDTH || newX-int32(projectile.hitBoxRadius) < 0 {
			state.destroyProjectile(idx)
			return
		}
		if newY+int32(projectile.hitBoxRadius) > WINDOW_HEIGHT || newY-int32(projectile.hitBoxRadius) < 0 {
			state.destroyProjectile(idx)
			return
		}
		projectile.x = newX
		projectile.y = newY
		projectile.rect.X += int32(projectile.scaler[0])
		projectile.rect.Y += int32(projectile.scaler[1])
		(*state.Projectiles)[idx] = projectile
	}
}

func (state *gameState) destroyEnemy(enemy *Enemy) {
	for idx, obj := range *state.Enemies {
		if obj.id == enemy.id && enemy.id != 0 {
			if idx == 0 {
				if len(*state.Enemies) > 1 {
					(*state.Enemies) = (*state.Enemies)[1:]
					return
				}
				state.Enemies = &[]Enemy{}
				return
			}
			(*state.Enemies)[idx] = (*state.Enemies)[len(*state.Enemies)-1]
			(*state.Enemies) = (*state.Enemies)[:len(*state.Enemies)-1]
		}
	}
}

func (state *gameState) destroyProjectile(idx int) {
	if idx == 0 {
		if len(*state.Projectiles) > 1 {
			(*state.Projectiles) = (*state.Projectiles)[1:]
			return
		}
		state.Projectiles = &[]Projectile{}
		return
	}
	(*state.Projectiles)[idx] = (*state.Projectiles)[len(*state.Projectiles)-1]
	(*state.Projectiles) = (*state.Projectiles)[:len(*state.Projectiles)-1]
}

func willCollide(state *gameState, x, y int32, hitBoxRadius uint8) (bool, *Enemy) {
	for _, enemy := range *state.Enemies {
		distance := math.Sqrt(math.Pow(float64(enemy.x-x), 2) + math.Pow(float64(enemy.y-y), 2))
		if distance < float64(enemy.hitBoxRadius+hitBoxRadius) {
			return true, &enemy
		}
	}
	return false, &Enemy{}
}
