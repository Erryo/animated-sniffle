package main

func initPlayer() *player {
	player := player{}
	player.sprite = initSprite(0, 0, 2)
	player.hp = 8
	player.vector = [2]int32{0, 0}
	return &player
}
