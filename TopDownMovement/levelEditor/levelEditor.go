package leveleditor

import (
	"fmt"
	"log"
	"strconv"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type state struct {
	window       *sdl.Window
	renderer     *sdl.Renderer
	textureAtlas *sdl.Texture
	font         *ttf.Font
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
	FontSize                  = 32
	TextureAtlasW             = 6
	TextureAtlasH             = 9
	CameraStep                = 14
)

func StartLevelEditor() {
	if err := ttf.Init(); err != nil {
		log.Panicln(err)
	}
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

	state.font, err = ttf.OpenFont("assets/font/ka1.ttf", FontSize)
	if err != nil {
		log.Panicln(err)
	}

	state.aosTiles = append(state.aosTiles, tile{x: 0, y: 0, typ: 5})
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

func (s *state) drawUI() {
	camXStr := strconv.Itoa(int(s.camera.X))
	camYStr := strconv.Itoa(int(s.camera.Y))
	coordText := "X:" + camXStr + "|| Y:" + camYStr
	//!NOTE Optimisation possible keep track of texture and check if text changd
	// if not simply use old texture
	coordTexture := s.createTextTexture(coordText)
	s.renderer.Copy(coordTexture, nil, &sdl.Rect{X: 0, Y: 0, W: int32(len(coordText)) * FontSize / 2, H: FontSize})
	coordTexture.Destroy()
}

func (s *state) createTextTexture(str string) *sdl.Texture {
	surf, err := s.font.RenderUTF8Solid(str, sdl.Color{255, 0, 0, 0})
	if err != nil {
		log.Panicln(err)
	}
	defer surf.Free()
	texture, err := s.renderer.CreateTextureFromSurface(surf)
	if err != nil {
		log.Panicln(err)
	}
	return texture
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
			case *sdl.MouseButtonEvent:
				s.doLMBPress(e)

			}
		}
		s.renderer.SetDrawColor(255, 255, 255, 255)
		s.renderer.Clear()
		s.drawGrid()
		s.drawTiles()
		s.drawUI()

		s.renderer.Present()
	}
}

func (s *state) doLMBPress(e *sdl.MouseButtonEvent) {
	var x, y int32
	x, y = e.X, e.Y
	switch e.Button {
	case 1:
		fmt.Println("mouse L")
		s.getCellGlobalPos(x, y)
	case 2:
		fmt.Println("mouse M")

	}
}

func (s *state) getCellGlobalPos(x, y int32) {
	fmt.Println(x, y)
	// !NOTE Need to change coords to fit top left corner of tile
	tile := tile{x: x - s.camera.X, y: y - s.camera.Y, typ: 1}
	fmt.Println(tile.x, tile.y)
	s.aosTiles = append(s.aosTiles, tile)
	fmt.Println(s.aosTiles)
}

func (s *state) drawTiles() {
	_, wH := s.window.GetSize()
	s.renderer.FillRect(&sdl.Rect{X: ArtSize*12 + s.camera.X, Y: wH + s.camera.Y, W: 16, H: 16})
	s.renderer.SetDrawColor(0, 0, 0, 255)
	src := sdl.Rect{W: ArtSize, H: ArtSize}
	dst := sdl.Rect{W: ArtSize, H: ArtSize}
	for _, tile := range s.aosTiles {
		srcX, srcY := indexToPosition(tile.typ)
		src.X = int32(srcX) * ArtSize
		src.Y = int32(srcY) * ArtSize

		dst.X = tile.x + s.camera.X
		dst.Y = tile.y + s.camera.Y

		s.renderer.Copy(s.textureAtlas, &src, &dst)
	}
}

func indexToPosition(idx uint8) (uint8, uint8) {
	x := (idx) % TextureAtlasW
	y := (idx) / TextureAtlasW
	return x, y
}
