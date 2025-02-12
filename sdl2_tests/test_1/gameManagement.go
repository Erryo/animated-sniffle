package main

import (
	"math"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (state *gameState) gameLoop() {
	running := true

outerGameLoop:
	for running {

		if state.doInput() {
			running = false
			break outerGameLoop
		}

		state.prepareScene()
		state.Update()
		state.drawAllObjects()

		sdl.Delay(GAME_UPDATE_DELAY)
	}
}

func (state *gameState) loadMedia() {
	var err error
	if state.backgroundImage, err = img.LoadTexture(state.renderer, "media/background.png"); err != nil {
		panic(err)
	}
}

// returns true if recieved quit signal
func (state *gameState) doInput() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			state.QuitGame()
			state.CloseSDL()
			return true
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				state.doKeyDown(e)
			} else if e.Type == sdl.KEYUP {
				state.doKeyUp(e)
			}

		}
	}
	return false
}

func (state *gameState) doKeyDown(event *sdl.KeyboardEvent) {
	if event.Repeat != 0 {
		return
	}
	switch event.Keysym.Scancode {
	case sdl.SCANCODE_W:
		state.Player.eventList[0] = true

	case sdl.SCANCODE_A:
		state.Player.eventList[1] = true

	case sdl.SCANCODE_S:
		state.Player.eventList[2] = true

	case sdl.SCANCODE_D:
		state.Player.eventList[3] = true

	case sdl.SCANCODE_SPACE:
		state.Player.eventList[4] = true
	}
}

func (state *gameState) doKeyUp(event *sdl.KeyboardEvent) {
	switch event.Keysym.Scancode {
	case sdl.SCANCODE_W:
		state.Player.eventList[0] = false

	case sdl.SCANCODE_A:
		state.Player.eventList[1] = false
	case sdl.SCANCODE_S:
		state.Player.eventList[2] = false

	case sdl.SCANCODE_D:
		state.Player.eventList[3] = false
	case sdl.SCANCODE_SPACE:
		state.Player.eventList[4] = false

	}
}

func (state *gameState) QuitGame() {
	state.Player.texture.Destroy()
	state.Player.texture = nil
	state.backgroundImage.Destroy()
	state.backgroundImage = nil
	state.renderer.Destroy()
	state.renderer = nil
	state.window.Destroy()
	state.window = nil
}

func (state *gameState) CloseSDL() {
	img.Quit()
	sdl.Quit()
}

func (state *gameState) AssignID() uint16 {
	state.nextID++
	return state.nextID - 1
}

func (state *gameState) initObject(color [3]uint8, x, y, w, h int32, hp int8) {
	rect := sdl.Rect{X: x - w/2, Y: y - h/2, W: w, H: h}
	radius := math.Sqrt(math.Pow(float64(h/2), 2) + math.Pow(float64(w/2), 2))
	enemy := Enemy{x: x, y: y, rect: &rect, color: color, id: state.AssignID(), hitBoxRadius: uint8(radius), hp: hp}
	*state.Enemies = append(*state.Enemies, enemy)
}

func (state *gameState) initProjectile(color [3]uint8, lifeLength uint16, scaler [2]int16, damage uint8, x, y, w, h int32) {
	rect := sdl.Rect{X: x - w/2, Y: y - h/2, W: w, H: h}
	radius := math.Sqrt(math.Pow(float64(h/2), 2) + math.Pow(float64(w/2), 2))
	projectile := Projectile{x: x, id: state.AssignID(), y: y, rect: &rect, color: color, hitBoxRadius: uint8(radius), lifeLength: lifeLength, scaler: scaler, damage: damage}
	*state.Projectiles = append(*state.Projectiles, projectile)
}

// texturePath := ,,media/name.png
func (state *gameState) initPlayer(x, y int32, speed uint8, texturePath string) {
	texture, err := img.LoadTexture(state.renderer, texturePath)
	if err != nil {
		panic(err)
	}
	eventList := make([]bool, 6)

	_, _, w, h, err := texture.Query()
	if err != nil {
		panic(err)
	}

	radius := math.Sqrt(math.Pow(float64(h/2), 2) + math.Pow(float64(w/2), 2))
	player := Player{x: x, y: y, texture: texture, id: state.AssignID(), speed: speed, eventList: eventList, hitBoxRadius: uint8(radius), ammo: 2, cooldown: 0}
	state.Player = &player
}

func (state *gameState) Update() {
	state.Player.checkEventList(state)
	// needs Testing <- written in a hurry
	if state.Player.cooldown > 0 {
		state.Player.cooldown -= 1
		if state.Player.cooldown == 0 {
			state.Player.ammo = 12
		}
	}
	state.moveProjectiles()
	state.blit(*state.Player)
}
