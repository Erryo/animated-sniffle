package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

func (state *state) setRuneToCoord() {
	state.runeToCoord = &map[rune][2]uint8{
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

// NOTE: rename to drawString
func (state *state) print(renderer *sdl.Renderer, str string, scale uint8, x, y int32, color [3]uint8) {
	str = strings.ToLower(str)
	state.fontMap.SetColorMod(color[0], color[1], color[2])
	src := &sdl.Rect{X: 0, Y: 0, W: FONT_W, H: FONT_H}
	dst := &sdl.Rect{X: x, Y: y, W: FONT_W * int32(scale), H: FONT_H * int32(scale)}
	for _, rune := range str {
		src.X = int32((*state.runeToCoord)[rune][0] * FONT_W)
		src.Y = int32((*state.runeToCoord)[rune][1] * FONT_H)
		renderer.Copy(state.fontMap, src, dst)
		dst.X += int32(FONT_W*scale) - 4
		// just test
		src.X += int32(FONT_W * scale)
	}
	state.fontMap.SetColorMod(255, 255, 255)
}

func (state *state) drawElements() {
	var str string
	var dataStr string
	for _, elem := range *state.currentLevel.dataElements {
		switch data := elem.data.(type) {
		case *string:
			dataStr = *data
		case *int:
			fmt.Println("drawn int")
			dataStr = strconv.Itoa(*data)
		case *float64:
			dataStr = strconv.FormatFloat(*data, 'f', 3, 64)
		case *float32:
			dataStr = strconv.FormatFloat(float64(*data), 'f', 3, 32)
		case *uint:
			dataStr = strconv.Itoa(int(*data))
		case *uint8:
			dataStr = strconv.Itoa(int(*data))
		case *uint16:
			dataStr = strconv.Itoa(int(*data))
		case *uint32:
			dataStr = strconv.Itoa(int(*data))
		case *uint64:
			dataStr = strconv.Itoa(int(*data))
		case *int8:
			dataStr = strconv.Itoa(int(*data))
		case *int16:
			dataStr = strconv.Itoa(int(*data))
		case *int32:
			dataStr = strconv.Itoa(int(*data))
		case *int64:
			dataStr = strconv.Itoa(int(*data))
		case nil:
		default:
			fmt.Printf("%T", elem.data)
		}
		str = elem.prefix + dataStr
		state.print(state.renderer, str, elem.size, elem.x, elem.y, elem.color)
	}
}

func (lvl *level) addElement(data interface{}, name, prefix string, x, y int32, size uint8, color [3]uint8, id uint16) error {
	if name != "" {
		for idx, elem := range *lvl.dataElements {
			if elem.name == name {
				return fmt.Errorf("element with name %v already exists at: %v", name, &(*lvl.dataElements)[idx])
			}
		}
	}
	element := dataElement{data: data, name: name, prefix: prefix, x: x, y: y, size: size, color: color, id: id}
	*lvl.dataElements = append(*lvl.dataElements, element)
	return nil
}

func (level *level) getElementByData(data interface{}) *dataElement {
	for idx, elem := range *level.dataElements {
		if elem.data == data {
			return &(*level.dataElements)[idx]
		}
	}
	return nil
}

func (lvl *level) getElementByName(name string) (*dataElement, error) {
	if name == "" {
		return nil, fmt.Errorf("Cant get elemnt by empty name")
	}

	for idx, elem := range *lvl.dataElements {
		fmt.Println(elem.name)
		if elem.name == name {
			return &(*lvl.dataElements)[idx], nil
		}
	}
	return nil, fmt.Errorf("No element found")
}

func (lvl *level) clearElements() {
	lvl.dataElements = &[]dataElement{}
}
