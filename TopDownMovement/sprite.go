package main

import "github.com/veandco/go-sdl2/sdl"

func initSprite(textureIdx, sType uint8, z int) sprite {
	sprite := sprite{textureIndex: textureIdx, sType: sType, zLevel: z}
	sprite.boundingBox = sdl.Rect{X: 0, Y: 0, W: ArtSize, H: ArtSize}
	sprite.textureX = 0 - sprite.boundingBox.W/2
	sprite.textureY = 0 - sprite.boundingBox.H/2
	return sprite
}
