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
}

func (player *Player) moveUp(state *gameState, speedModifier float32) {
	newPlayerCoordinate := player.y - int32(float32(player.speed)*speedModifier)
	if newPlayerCoordinate-int32(player.hitBoxRadius) > 0 {
		if willColide, _ := willCollide(state, player.x, newPlayerCoordinate, player.hitBoxRadius); !willColide {
			player.y = newPlayerCoordinate
			if player.rotation == 0 || player.rotation == 360 {
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
		if willColide, _ := willCollide(state, player.x, newPlayerCoordinate, player.hitBoxRadius); !willColide {
			player.y = newPlayerCoordinate
			if player.rotation == 180 {
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
		if willColide, _ := willCollide(state, newPlayerCoordinate, player.y, player.hitBoxRadius); !willColide {
			player.x = newPlayerCoordinate
			if player.rotation == 270 {
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
		if willColide, _ := willCollide(state, newPlayerCoordinate, player.y, player.hitBoxRadius); !willColide {
			player.x = newPlayerCoordinate
			if player.rotation == 90 {
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
		player.cooldown = 64 * 2
	}
	state.initProjectile(RED, 64*3, [2]int16{12, 0}, 1, player.x, player.y, 10, 10)
	player.ammo -= 1
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
