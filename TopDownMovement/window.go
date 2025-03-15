package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func (s *state) prepareScene() {
	s.renderer.SetDrawColor(255, 255, 255, 255)
	// s.renderer.SetDrawColor(0, 0, 0, 0)
	s.renderer.Clear()
}

func (s *state) drawTiles() {
	wW, wH := s.window.GetSize()

	newTileSize := wH / MaxViewHeigth
	fmt.Println(newTileSize)
	mapStartX := (wW - newTileSize*MaxViewWidth) / 2
	mapStartX = 0

	var srcX, srcY uint8
	src := sdl.Rect{W: ArtSize, H: ArtSize}
	dst := sdl.Rect{W: newTileSize, H: newTileSize}

	for y, row := range s.currentLevel.tiles {
	inner:
		for x, tile := range row {
			if tile == EmptyTile {
				continue inner
			}

			dst.X = int32(x)*newTileSize + s.currentLevel.camera.X + mapStartX
			dst.Y = int32(y)*newTileSize + s.currentLevel.camera.Y

			srcX, srcY = indexToPosition(tile)
			src.X = int32(srcX) * ArtSize
			src.Y = int32(srcY) * ArtSize

			s.renderer.Copy(s.textureAtlas, &src, &dst)
		}
	}
}

func indexToPosition(idx uint8) (uint8, uint8) {
	x := (idx) % TextureAtlasW
	y := (idx) / TextureAtlasW
	return x, y
}
