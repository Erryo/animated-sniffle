package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func (s *state) prepareScene() {
	// s.renderer.SetDrawColor(0, 0, 0, 0)

	s.windowW, s.windowH = s.window.GetSize()
	s.window.SetSize(s.windowH, s.windowH)

	s.updateRendererScale()

	s.renderer.SetDrawColor(255, 255, 255, 255)
	s.renderer.Clear()
}

func (s *state) drawTiles() {
	var srcX, srcY uint8

	src := sdl.Rect{W: ArtSize, H: ArtSize}
	dst := sdl.Rect{W: ArtSize, H: ArtSize}

	if s.currentLevel.camera.Y < 0 {
		s.currentLevel.camera.Y = 0
	}
	if s.currentLevel.camera.X < 0 {
		s.currentLevel.camera.X = 0
	}
	if s.currentLevel.camera.Y+MaxViewHeigth > MaxLevelHeigth {
		s.currentLevel.camera.Y = MaxLevelHeigth - MaxViewHeigth
	}
	if s.currentLevel.camera.X+MaxViewWidth > MaxLevelWidth {
		s.currentLevel.camera.X = MaxLevelWidth - MaxViewWidth
	}

	for y := 0 + s.currentLevel.camera.Y; y < MaxViewHeigth+s.currentLevel.camera.Y; y++ {
	inner:
		for x := 0 + s.currentLevel.camera.X; x < MaxViewWidth+s.currentLevel.camera.X; x++ {
			tile := s.currentLevel.tiles[y+s.currentLevel.camera.Y][x+s.currentLevel.camera.X]
			if tile == EmptyTile {
				continue inner
			}

			dst.X = int32(x)*ArtSize + s.currentLevel.camera.X
			dst.Y = int32(y)*ArtSize + s.currentLevel.camera.Y

			srcX, srcY = indexToPosition(tile)
			src.X = int32(srcX) * ArtSize
			src.Y = int32(srcY) * ArtSize

			s.renderer.Copy(s.textureAtlas, &src, &dst)
		}
	}
}

func (s *state) updateRendererScale() {
	scaleX := (float32(s.windowH) / MaxViewHeigth) / ArtSize
	s.renderer.SetScale(scaleX, scaleX)
}

func indexToPosition(idx uint8) (uint8, uint8) {
	x := (idx) % TextureAtlasW
	y := (idx) / TextureAtlasW
	return x, y
}
