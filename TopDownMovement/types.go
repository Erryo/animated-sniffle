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
	fontAtlas    *sdl.Texture
	textureAtlas *sdl.Texture
	runeToCoord  *map[rune][2]uint8

	// time difference between last frame and the frame before the last frame in MS
	deltaTime float32
}

type level struct {
	// unique
	name            string
	music           *mix.Music
	backgroundImage *sdl.Texture
	// all elements that are not tiles
	// sorted by z(depth) order
	blitables [NumberOfZLevels]([]bliter)
	player    *player
	nextID    uint16
}

type projectile struct {
	sprite
	movable
}

type player struct {
	sprite
	movable
	killable
}
type enemy struct {
	movable
	sprite
	killable
}

type killable struct {
	hp uint16
}

type movable struct {
	vector [2]int32
}

type sprite struct {
	texture *sdl.Texture
	// -1 for top aka ui
	zLevel int
	// used to diferenciate
	sType       uint8
	boundingBox sdl.Rect
	// the coords of the Top Left Corner of the texture
	textureX, textureY int32
}

type textUIElement struct {
	// used no get element
	// Tag does not have to be unique
	// id has to be
	tag    string
	id     uint16
	prefix string
	data   interface{}
	suffix string
	x, y   int32
	size   uint8
	color  [3]uint8
}

// Pointer to sprite is a bliter
// but sprite is not
// Bliter must be a pointer to ...
type bliter interface {
	blit()
	destroy()
}

const (
	WindowWidth         = 1200
	WindowHeigth        = 900
	NumberOfZLevels     = 4
	PathToMainMenuMusic = "assets/music/mainMenu.ogg"
	PathToTextureAtlas  = "assets/png/textureAtlas.png"
	PathToFontAtlas     = "assets/png/fontAtlas.png"
)
