package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Better called level state
type gameState struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	Enemies         *[]Enemy
	Projectiles     *[]Projectile
	Player          *Player
	nextID          uint16
	backgroundImage *sdl.Texture
}

type Projectile struct {
	damage       uint8
	lifeLength   uint16
	hitBoxRadius uint8
	rect         *sdl.Rect
	color        [3]uint8
	// By how much the position sholud change each frame
	scaler [2]int16
	id     uint16
	x, y   int32
}
type Enemy struct {
	id           uint16
	x, y         int32
	hp           int8
	rect         *sdl.Rect
	hitBoxRadius uint8
	direction    int8
	color        [3]uint8
}
type Player struct {
	id    uint16
	speed uint8
	ammo  uint8
	// shooting cooldown
	cooldown     uint8
	x, y         int32
	texture      *sdl.Texture
	hitBoxRadius uint8
	rotation     int16
	// The eventList tells if a key was pressed down and not lifted up
	// the order: moveUp moveL moveDown moveRight Fire
	eventList []bool
}

const (
	WINDOW_WIDTH  = 900
	WINDOW_HEIGHT = 600
	//  time/frequency : 1000ms/24fps
	GAME_UPDATE_DELAY = 17
)

// Colors
var (
	MAGENTA = [3]uint8{231, 0, 106}
	ORANGE  = [3]uint8{243, 152, 1}
	YELLOW  = [3]uint8{248, 248, 69}
	BLUE    = [3]uint8{1, 104, 183}
	CYAN    = [3]uint8{50, 103, 183}
	RED     = [3]uint8{255, 0, 0}
)
