package main

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type state struct {
	window       *sdl.Window
	renderer     *sdl.Renderer
	levels       []*level
	currentLevel *level
	fontMap      *sdl.Texture
	runeToCoord  *map[rune][2]uint8

	// time difference between last frame and the frame before the last frame in MS
	deltaTime float32
}

type level struct {
	name            string
	music           *mix.Music
	dataElements    *[]dataElement
	enemies         *[]enemy
	projectiles     *[]projectile
	player          *player
	backgroundImage *sdl.Texture
	nextID          uint16
}

type projectile struct {
	damage       uint8
	lifeLength   uint16
	hitBoxRadius uint8
	rect         *sdl.Rect
	color        [3]uint8
	speed        uint8
	vector       [2]float32
	id           uint16
	x, y         int32
}

type enemy struct {
	id           uint16
	x, y         int32
	vector       [2]int16
	hp           int8
	rect         *sdl.Rect
	hitBoxRadius uint8
	color        [3]uint8
}

type player struct {
	id           uint16
	texture      *sdl.Texture
	hitBoxRadius uint8
	x, y         int32
	rotation     int16
	speed        uint8
	vector       [2]int8
	// Will act as a scaler to the vector, it wont be multiplied
	// rather divided To avoid the introduction of floats
	// The bigger the momentum the slower the player
	//	momentum      uint8
	cooldown      uint16
	magazine_size uint8
	ammo          uint8
	// The eventList tells if a key was pressed down and not lifted up
	// the order: moveUp moveL moveDown moveRight Fire RotateRight RotateLeft
	eventList []bool
	reloading bool
	shootEff  *mix.Chunk
}

type dataElement struct {
	data   interface{}
	name   string
	prefix string
	suffix string
	x, y   int32
	size   uint8
	color  [3]uint8
	id     uint16
}

const (
	WINDOW_WIDTH  = 900
	WINDOW_HEIGHT = 600
	//  time/frequency : 1000ms/60fps = 17
	GAME_UPDATE_DELAY      = 1000 / 50
	PLAYER_MAG_SIZE        = 10
	PLAYER_RELOAD_COOLDOWD = 60 * 2
	PLAYER_SPEED           = 8
	FONT_MAP_W             = 5
	FONT_MAP_H             = 5
	FONT_H                 = 20
	FONT_W                 = 20
)

// Colors
var (
	MAGENTA    = [3]uint8{231, 0, 106}
	ORANGE     = [3]uint8{243, 152, 1}
	YELLOW     = [3]uint8{248, 248, 69}
	BLUE       = [3]uint8{1, 104, 183}
	CYAN       = [3]uint8{50, 103, 183}
	RED        = [3]uint8{255, 0, 0}
	WHITE      = [3]uint8{255, 255, 255}
	ALL_COLORS = [][3]uint8{MAGENTA, ORANGE, YELLOW, BLUE, CYAN}
)
