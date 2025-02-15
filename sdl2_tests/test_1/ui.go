package main

import (
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

func (textMan *TextManager) setDict() {
	textMan.dict = &map[rune][2]uint8{
		'a':  {0, 0},
		'b':  {1, 0},
		'c':  {2, 0},
		'd':  {3, 0},
		'e':  {4, 0},
		'f':  {5, 0},
		'g':  {6, 0},
		'h':  {7, 0},
		'i':  {0, 1},
		'j':  {1, 1},
		'k':  {2, 1},
		'l':  {3, 1},
		'm':  {4, 1},
		'n':  {5, 1},
		'o':  {6, 1},
		'p':  {7, 1},
		'q':  {0, 2},
		'r':  {1, 2},
		's':  {2, 2},
		't':  {3, 2},
		'u':  {4, 2},
		'v':  {5, 2},
		'w':  {6, 2},
		'x':  {7, 2},
		'y':  {0, 3},
		'z':  {1, 3},
		'0':  {2, 3},
		'1':  {3, 3},
		'2':  {4, 3},
		'3':  {5, 3},
		'4':  {6, 3},
		'5':  {7, 3},
		'6':  {0, 4},
		'7':  {1, 4},
		'8':  {2, 4},
		'9':  {3, 4},
		':':  {4, 4},
		',':  {5, 4},
		'.':  {6, 4},
		'\\': {7, 4},
		' ':  {0, 5},
	}
}

func (textMan *TextManager) print(renderer *sdl.Renderer, str string, scale uint8, x, y int32, r, g, b uint8) {
	str = strings.ToLower(str)
	textMan.fontMap.SetColorMod(r, g, b)
	src := &sdl.Rect{X: 0, Y: 0, W: FONT_W, H: FONT_H}
	dst := &sdl.Rect{X: x, Y: y, W: FONT_W * int32(scale), H: FONT_H * int32(scale)}
	for _, rune := range str {
		src.X = int32((*textMan.dict)[rune][0] * FONT_W)
		src.Y = int32((*textMan.dict)[rune][1] * FONT_H)
		renderer.Copy(textMan.fontMap, src, dst)
		dst.X += int32(FONT_W*scale) - 4
		// just test
		src.X += int32(FONT_W * scale)
	}
	textMan.fontMap.SetColorMod(255, 255, 255)
}
