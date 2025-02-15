package main

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

// Better called level state
type gameState struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	music           *mix.Music
	Enemies         *[]Enemy
	Projectiles     *[]Projectile
	Player          *Player
	nextID          uint16
	backgroundImage *sdl.Texture
	TextManager     *TextManager
}

type Projectile struct {
	damage       uint8
	lifeLength   uint16
	hitBoxRadius uint8
	rect         *sdl.Rect
	color        [3]uint8
	speed        uint8
	scaler       [2]float32
	id           uint16
	x, y         int32
}
type Enemy struct {
	id           uint16
	x, y         int32
	scaler       [2]int16
	hp           int8
	rect         *sdl.Rect
	hitBoxRadius uint8
	color        [3]uint8
}
type Player struct {
	id            uint16
	shootEff      *mix.Chunk
	speed         uint8
	magazine_size uint8
	ammo          uint8
	cooldown      uint16
	x, y          int32
	texture       *sdl.Texture
	hitBoxRadius  uint8
	rotation      int16
	// The eventList tells if a key was pressed down and not lifted up
	// the order: moveUp moveL moveDown moveRight Fire RotateRight RotateLeft
	eventList []bool
	reloading bool
}
type TextManager struct {
	fontMap *sdl.Texture
	dict    *map[rune][2]uint8
}

const (
	WINDOW_WIDTH  = 900
	WINDOW_HEIGHT = 600
	//  time/frequency : 1000ms/60fps
	GAME_UPDATE_DELAY      = 17
	PLAYER_MAG_SIZE        = 10
	PLAYER_RELOAD_COOLDOWD = 60 * 2
	FONT_MAP_W             = 8
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
	ALL_COLORS = [][3]uint8{MAGENTA, ORANGE, YELLOW, BLUE, CYAN}
)
