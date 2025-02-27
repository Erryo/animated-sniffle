package main

import "github.com/veandco/go-sdl2/sdl"

func (player *player) checkEventList(lvl *level) {
	if player.eventList[0] {
		if player.eventList[1] || player.eventList[3] {
			player.moveUp(1 / 1.4142)
		} else {
			player.moveUp(1)
		}
	}
	if player.eventList[1] {
		if player.eventList[0] || player.eventList[2] {
			player.moveLeft(1 / 1.4142)
		} else {
			player.moveLeft(1)
		}
	}
	if player.eventList[2] {
		// To make diagonal movement the same speed as the normal movement
		if player.eventList[1] || player.eventList[3] {
			player.moveDown(1 / 1.4142)
		} else {
			player.moveDown(1)
		}
	}
	if player.eventList[3] {
		if player.eventList[0] || player.eventList[2] {
			player.moveRight(1 / 1.4142)
		} else {
			player.moveRight(1)
		}
	}
	if player.eventList[4] {
		player.fire(lvl)
	}
	if player.eventList[5] {
		player.changeRotation(5)
	}
	if player.eventList[6] {
		player.changeRotation(-5)
	}
}

func (player *player) reduceVector() {
	var playerSpeedDecrease int8
	if sdl.GetPerformanceCounter()%3 == 0 {
		playerSpeedDecrease = 1
	}
	if player.vector[0] < 0 {
		player.vector[0] = player.vector[0] + playerSpeedDecrease
	}
	if player.vector[0] > 0 {
		player.vector[0] = player.vector[0] - playerSpeedDecrease
	}

	if player.vector[1] < 0 {
		player.vector[1] = player.vector[1] + playerSpeedDecrease
	}
	if player.vector[1] > 0 {
		player.vector[1] = player.vector[1] - playerSpeedDecrease
	}
}

func (player *player) movePlayer(lvl *level) {
	newX := player.x + int32(player.vector[0])
	newY := player.y + int32(player.vector[1])

	if newY+int32(player.hitBoxRadius) > WINDOW_HEIGHT || newY-int32(player.hitBoxRadius) < 0 {
		newY = player.y
	}
	if newX+int32(player.hitBoxRadius) > WINDOW_WIDTH || newX-int32(player.hitBoxRadius) < 0 {
		newX = player.x
	}

	if willColide, _ := lvl.willCollide(newX, newY, player.hitBoxRadius, -1); willColide {
		return
	}

	player.y = newY
	player.x = newX
}

func (player *player) moveUp(speedModifier float32) {
	player.vector[1] = -int8(float32(player.speed) * speedModifier)
	if player.rotation == 0 || player.rotation == 360 || (player.rotation > 350 || player.rotation < 10) {
		player.rotation = 0
		return
	}
	if player.rotation >= 180 {
		player.changeRotation(15)
	} else if player.rotation < 180 {
		player.changeRotation(-15)
	}
}

func (player *player) moveDown(speedModifier float32) {
	player.vector[1] = int8(float32(player.speed) * speedModifier)
	if player.rotation == 180 || (player.rotation > 170 && player.rotation < 190) {
		player.rotation = 180
		return
	}
	if player.rotation >= 0 && player.rotation < 180 {
		player.changeRotation(15)
	} else if player.rotation <= 360 && player.rotation > 180 {
		player.changeRotation(-15)
	}
}

func (player *player) moveLeft(speedModifier float32) {
	player.vector[0] = -int8(float32(player.speed) * speedModifier)
	if player.rotation == 270 || (player.rotation > 260 && player.rotation < 280) {
		player.rotation = 270
		return
	}
	if player.rotation >= 90 && player.rotation < 270 {
		player.changeRotation(15)
	} else if player.rotation < 90 || player.rotation <= 360 {
		player.changeRotation(-15)
	}
}

func (player *player) moveRight(speedModifier float32) {
	player.vector[0] = int8(float32(player.speed) * speedModifier)

	if player.rotation == 90 || (player.rotation > 80 && player.rotation < 100) {
		player.rotation = 90
		return
	}
	if player.rotation < 270 && player.rotation > 90 {
		player.changeRotation(-15)
	} else if player.rotation >= 270 || player.rotation < 90 {
		player.changeRotation(15)
	}
}

// needs Testing <- written in a hurry
func (player *player) fire(lvl *level) {
	if player.ammo <= 0 || player.cooldown != 0 {
		return
	}
	if player.ammo-1 == 0 {
		player.cooldown = PLAYER_RELOAD_COOLDOWD
		player.reloading = true
	} else {
		player.cooldown = 30
	}
	player.shootEff.Play(-1, 0)
	lvl.initProjectile(RED, 64*3, 10, 1, player.x, player.y, 10, 10)
	player.ammo -= 1
}

func (player *player) handleFireCooldown() {
	if player.cooldown > 0 {
		player.cooldown -= 1
		if player.cooldown == 0 {
			if player.reloading {
				player.ammo = player.magazine_size
				player.reloading = false
			}
		}
	}
}

func (player *player) changeRotation(angle int16) {
	newAngle := player.rotation + angle
	if newAngle < 0 {
		player.rotation = 360 + newAngle
	} else if newAngle >= 360 {
		player.rotation = newAngle - 360
	} else {
		player.rotation += angle
	}
}
