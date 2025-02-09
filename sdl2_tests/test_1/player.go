package main

import (
	"fmt"
	"math"
)

func (player *Player) checkEventList(state *gameState) {
	if player.eventList[0] {
		player.moveUp(state)
	}
	if player.eventList[1] {
		player.moveLeft(state)
	}
	if player.eventList[2] {
		player.moveDown(state)
	}
	if player.eventList[3] {
		player.moveRight(state)
	}
	if player.eventList[4] {
		player.fire()
	}
}

func (player *Player) moveUp(state *gameState) {
	if player.y-int32(player.speed)-int32(player.hitBoxRadius) > 0 && !player.willCollide(state, player.x, player.y-int32(player.speed)) {
		player.y += -int32(player.speed)
		if player.rotation == 0 || player.rotation == 360 {
			return
		}
		if player.rotation >= 180 {
			player.changeRotation(15)
		}
		if player.rotation < 180 {
			player.changeRotation(-15)
		}
	}
}

func (player *Player) moveDown(state *gameState) {
	if player.y+int32(player.speed)+int32(player.hitBoxRadius) < WINDOW_HEIGHT && !player.willCollide(state, player.x, player.y+int32(player.speed)) {
		player.y += int32(player.speed)
		if player.rotation == 180 {
			return
		}
		if player.rotation >= 0 {
			player.changeRotation(15)
		} else if player.rotation <= 360 {
			player.changeRotation(-15)
		}
	}
}

func (player *Player) moveLeft(state *gameState) {
	if player.x-int32(player.speed)-int32(player.hitBoxRadius) > 0 && !player.willCollide(state, player.x-int32(player.speed), player.y) {
		player.x += -int32(player.speed)
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

func (player *Player) moveRight(state *gameState) {
	if player.x+int32(player.speed)+int32(player.hitBoxRadius) < WINDOW_WIDTH && !player.willCollide(state, player.x+int32(player.speed), player.y) {
		player.x += int32(player.speed)
		if player.rotation == 90 {
			return
		}
		if player.rotation < 270 && player.rotation > 90 {
			player.changeRotation(-15)
		}
		if player.rotation >= 270 || player.rotation < 90 {
			player.changeRotation(15)
		}

	}
}

func (player *Player) fire() {
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
	fmt.Println(player.rotation, newAngle)
}

func (player *Player) willCollide(state *gameState, x, y int32) bool {
	for _, enemy := range *state.Enemies {
		distance := math.Sqrt(math.Pow(float64(enemy.x-x), 2) + math.Pow(float64(enemy.y-y), 2))
		if distance < float64(enemy.hitBoxRadius+player.hitBoxRadius) {
			return true
		}
	}
	return false
}
