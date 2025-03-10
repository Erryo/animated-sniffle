package main

func (player *Player) checkEventList(state *gameState) {
	if player.eventList[0] {
		if player.eventList[1] || player.eventList[3] {
			player.moveUp(state, 1/1.4142)
		} else {
			player.moveUp(state, 1)
		}
	}
	if player.eventList[1] {
		player.moveLeft(state, 1)
	}
	if player.eventList[2] {
		// To make diagonal movement the same speed as the normal movement
		if player.eventList[1] || player.eventList[3] {
			player.moveDown(state, 1/1.4142)
		} else {
			player.moveDown(state, 1)
		}
	}
	if player.eventList[3] {
		player.moveRight(state, 1)
	}
	if player.eventList[4] {
		player.fire(state)
	}
	if player.eventList[5] {
		player.changeRotation(5)
	}
	if player.eventList[6] {
		player.changeRotation(-5)
	}
}

func (player *Player) moveUp(state *gameState, speedModifier float32) {
	newPlayerCoordinate := player.y - int32(float32(player.speed)*speedModifier)
	if newPlayerCoordinate-int32(player.hitBoxRadius) > 0 {
		if willColide, _ := state.willCollide(player.x, newPlayerCoordinate, player.hitBoxRadius, -1); !willColide {
			player.y = newPlayerCoordinate
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
	}
}

func (player *Player) moveDown(state *gameState, speedModifier float32) {
	newPlayerCoordinate := player.y + int32(float32(player.speed)*speedModifier)
	if newPlayerCoordinate+int32(player.hitBoxRadius) < WINDOW_HEIGHT {
		if willColide, _ := state.willCollide(player.x, newPlayerCoordinate, player.hitBoxRadius, -1); !willColide {
			player.y = newPlayerCoordinate
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
	}
}

func (player *Player) moveLeft(state *gameState, speedModifier float32) {
	newPlayerCoordinate := player.x - int32(float32(player.speed)*speedModifier)
	if newPlayerCoordinate-int32(player.hitBoxRadius) > 0 {
		if willColide, _ := state.willCollide(newPlayerCoordinate, player.y, player.hitBoxRadius, -1); !willColide {
			player.x = newPlayerCoordinate
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
	}
}

func (player *Player) moveRight(state *gameState, speedModifier float32) {
	newPlayerCoordinate := player.x + int32(float32(player.speed)*speedModifier)
	if newPlayerCoordinate+int32(player.hitBoxRadius) < WINDOW_WIDTH {
		if willColide, _ := state.willCollide(newPlayerCoordinate, player.y, player.hitBoxRadius, -1); !willColide {
			player.x = newPlayerCoordinate
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
	}
}

// needs Testing <- written in a hurry
func (player *Player) fire(state *gameState) {
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
	state.initProjectile(RED, 64*3, 10, 1, player.x, player.y, 10, 10)
	player.ammo -= 1
}

func (player *Player) handleFireCooldown() {
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

func (player *Player) changeRotation(angle int16) {
	newAngle := player.rotation + angle
	if newAngle < 0 {
		player.rotation = 360 + newAngle
	} else if newAngle >= 360 {
		player.rotation = newAngle - 360
	} else {
		player.rotation += angle
	}
}
