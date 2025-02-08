package main

func (player *Player) checkEventList() {
	if player.eventList[0] {
		player.moveUp()
	}
	if player.eventList[1] {
		player.moveLeft()
	}
	if player.eventList[2] {
		player.moveDown()
	}
	if player.eventList[3] {
		player.moveRight()
	}
	if player.eventList[4] {
		player.fire()
	}
}

func (player *Player) moveUp() {
	if player.y-int(player.speed) > 0+20 {
		player.y += -int(player.speed)
	}
}

func (player *Player) moveDown() {
	if player.y+int(player.speed) < WINDOW_HEIGHT-80 {
		player.y += int(player.speed)
	}
}

func (player *Player) moveLeft() {
	if player.x-int(player.speed) > 0+20 {
		player.x += -int(player.speed)
	}
}

func (player *Player) moveRight() {
	if player.x+int(player.speed) < WINDOW_WIDTH-80 {
		player.x += int(player.speed)
	}
}

func (player *Player) fire() {
}
