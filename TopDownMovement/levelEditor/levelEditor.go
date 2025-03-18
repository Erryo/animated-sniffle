package leveleditor

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type state struct {
	window       *sdl.Window
	renderer     *sdl.Renderer
	textureAtlas *sdl.Texture
	tiles        [MaxLevelHeigth][MaxLevelWidth]uint8
	aosTiles     []tile
	camera       sdl.Point
}
type tile struct {
	x, y int32
	typ  uint8
}

const (
	WindowWidth, WindowHeigth = 900, 900
	MaxLevelHeigth            = 100
	MaxLevelWidth             = 100
	ArtSize                   = 16
	TextureAtlasW             = 6
	TextureAtlasH             = 9
	CameraStep                = 14
)

func StartLevelEditor() {
	window, err := sdl.CreateWindow("level editor", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		WindowWidth, WindowHeigth, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Panicln(err)
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Panicln(err)
	}

	atlas, err := img.LoadTexture(renderer, "assets/png/textureAtlas.png")
	if err != nil {
		log.Panicln("error loading textuer Atlas:", err)
	}

	state := state{renderer: renderer, window: window, textureAtlas: atlas}
	state.gameLoop()
}

func (s *state) drawGrid() {
	s.renderer.SetDrawColor(0, 0, 0, 255)
	wW, wH := s.window.GetSize()
	// ensures the grid ends at a multiple of art size smaller than window
	lineEndY := (wH / ArtSize) * ArtSize
	lineEndX := (wW / ArtSize) * ArtSize
	var x, y int32
	// Horizontal
	for j := int32(0); j < wH; j += ArtSize {
		y = j + s.camera.Y
		for ; y > wH; y -= lineEndY {
		}
		for ; y < 0; y += lineEndY {
		}

		s.renderer.DrawLine(0, y, wW, y)
	}
	// Vertical
	for i := int32(0); i < wW; i += ArtSize {
		x = i + s.camera.X
		for ; x > wW; x -= lineEndX {
		}
		for ; x < 0; x += lineEndX {
		}
		s.renderer.DrawLine(x, 0, x, wH)
	}
}

func (s *state) gameLoop() {
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("quitting")
				running = false
				return
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					switch e.Keysym.Scancode {
					case sdl.SCANCODE_W:
						s.camera.Y += CameraStep
					case sdl.SCANCODE_S:
						s.camera.Y -= CameraStep
					case sdl.SCANCODE_A:
						s.camera.X += CameraStep
					case sdl.SCANCODE_D:
						s.camera.X -= CameraStep
					}
				} else if e.Type == sdl.KEYUP {
				}

			}
		}
		s.renderer.SetDrawColor(255, 255, 255, 255)
		s.renderer.Clear()
		s.drawGrid()
		s.drawTiles()

		s.renderer.Present()
	}
}

func (s *state) drawTiles() {
	s.renderer.SetDrawColor(255, 0, 0, 255)
	_, wH := s.window.GetSize()

	s.renderer.FillRect(&sdl.Rect{X: ArtSize*12 + s.camera.X, Y: wH + s.camera.Y, W: 16, H: 16})
	s.renderer.SetDrawColor(0, 0, 0, 255)
	src := sdl.Rect{W: ArtSize, H: ArtSize}
	dst := sdl.Rect{W: ArtSize, H: ArtSize}
	for _, tile := range s.aosTiles {
		srcX, srcY := indexToPosition(tile.typ)
		src.X = int32(srcX) * ArtSize
		src.Y = int32(srcY) * ArtSize

		dst.X = int32(tile.x)*ArtSize + s.camera.X
		dst.Y = int32(tile.y)*ArtSize + s.camera.Y

		s.renderer.Copy(s.textureAtlas, &src, &dst)
	}
}

func indexToPosition(idx uint8) (uint8, uint8) {
	x := (idx) % TextureAtlasW
	y := (idx) / TextureAtlasW
	return x, y
}
