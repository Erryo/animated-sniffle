package main

import (
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
	}
}

func (player *Player) moveDown(state *gameState) {
	if player.y+int32(player.speed)+int32(player.hitBoxRadius) < WINDOW_HEIGHT && !player.willCollide(state, player.x, player.y+int32(player.speed)) {
		player.y += int32(player.speed)
	}
}

func (player *Player) moveLeft(state *gameState) {
	if player.x-int32(player.speed)-int32(player.hitBoxRadius) > 0 && !player.willCollide(state, player.x-int32(player.speed), player.y) {
		player.x += -int32(player.speed)
	}
}

func (player *Player) moveRight(state *gameState) {
	if player.x+int32(player.speed)+int32(player.hitBoxRadius) < WINDOW_WIDTH && !player.willCollide(state, player.x+int32(player.speed), player.y) {
		player.x += int32(player.speed)
	}
}

func (player *Player) fire() {
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
