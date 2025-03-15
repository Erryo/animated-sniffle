package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type Line struct {
	x1, x2, y1, y2 int32
}

func main() {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Shape Generator", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		800, 800, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	// var wg sync.WaitGroup
	// var mu sync.Mutex

	//	points := offsetPoints(createShape())
	//	points, lines := createTriangle()
	//	points := randomTriangle([]sdl.Point{})
	//	points0 := points
	//	points = offsetPoints(points)
	//	points0 = offsetPoints(points0)
	// points, _ = createTriangle()
	// fmt.Println("rt1")
	// fmt.Println("rt2")
	// fmt.Println("rt3")
	//	lines = offsetPointsOfLine(lines)
	//
	// points := offsetPoints(createTriangle())
	//
	// verts0 := pointsToVertex(points0)
	//	points := []sdl.Point{{414, 480}, {406, 417}, {468, 481}}
	//

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				sdl.Quit()
				return
				running = false
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					switch e.Keysym.Scancode {
					case sdl.SCANCODE_SPACE:
						sdl.Delay(1 * 1000)
					}
				}
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		renderer.SetDrawColor(255, 255, 255, 255)

		//	if err := renderer.RenderGeometry(nil, verts, nil); err != nil {
		//		log.Panic(err)
		//	}
		//	for range 10 {
		//		wg.Add(1)
		//		go func() {
		//			defer wg.Done()
		//			createAndDrawTriangle(renderer, window, &mu)
		//		}()
		//	}

		//	wg.Wait()
		createAndDrawTriangle(renderer, window)
		renderer.Present()

	}
	fmt.Println("Quit")
}

func createAndDrawTriangle(renderer *sdl.Renderer, window *sdl.Window) {
	var points []sdl.Point
	var verts []sdl.Vertex
	points = randomTriangle([]sdl.Point{})
	points = offsetPoints(points, window)
	verts = pointsToVertex(points)

	//	drawTriangle(pointsToLines(points), renderer)
	if err := renderer.RenderGeometry(nil, verts, nil); err != nil {
		log.Panic(err)
	}
}

func randomTriangle(points []sdl.Point) []sdl.Point {
	var radius int
	var point sdl.Point
	for radius <= 80 {
		radius = rand.Intn(350)
	}

	invalidPoint := true
	for len(points) < 3 {
		invalidPoint = true
	invPoint:
		for invalidPoint {
			point = sdl.Point{X: rand.Int31n(int32(radius)), Y: rand.Int31n(int32(radius))}
			for _, p := range points {
				if distBetwPoints(p, point) < float64(radius)/10 {
					fmt.Println(points, point, radius)
					continue invPoint
				}
			}
			invalidPoint = false
		}
		points = append(points, point)
	}
	fmt.Println(points)
	return points
}

func distBetwPoints(p1, p2 sdl.Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) + math.Pow(float64(p1.Y-p2.Y), 2))
}

func pointsToLines(points []sdl.Point) []Line {
	line := Line{points[0].X, points[1].X, points[0].Y, points[1].Y}
	line1 := Line{points[0].X, points[2].X, points[0].Y, points[2].Y}
	line2 := Line{points[1].X, points[2].X, points[1].Y, points[2].Y}
	return []Line{line, line1, line2}
}

func randInt(min, max int32) int32 {
	var num int32
	if min >= max {
		return max
	}
	for num <= min {
		num = rand.Int31n(max)
	}
	return num
}

func pointsToVertex(points []sdl.Point) []sdl.Vertex {
	var vertices []sdl.Vertex
	for _, p := range points {
		fpoint := sdl.FPoint{float32(p.X), float32(p.Y)}
		vert := sdl.Vertex{Position: fpoint, Color: sdl.Color{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}}
		vertices = append(vertices, vert)
	}
	return vertices
}

func drawTriangle(lines []Line, renderer *sdl.Renderer) {
	//lines = []Line{
	//	{814, 806, 880, 817},
	//	{814, 868, 880, 881},
	//	{806, 868, 817, 881},
	//}
	for _, line := range lines {
		err := renderer.DrawLine(line.x1, line.y1, line.x2, line.y2)
		fmt.Println(line)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func createTriangle() ([]sdl.Point, []Line) {
	points := []sdl.Point{}
	line := []Line{}
	p := sdl.Point{X: 100, Y: 0}
	p1 := sdl.Point{X: 0, Y: -100}
	p2 := sdl.Point{X: -100, Y: 0}
	points = append(points, p, p1, p2)

	line = append(line, Line{100, 0, 0, -100}, Line{100, -100, 0, 0}, Line{-100, 0, 0, -100})
	return points, line
}

func createShape() []sdl.Point {
	p1 := sdl.Point{X: 100, Y: 0}
	p := sdl.Point{X: 70, Y: 70}
	p2 := sdl.Point{X: 0, Y: 100}
	p7 := sdl.Point{X: 70, Y: -70}
	p4 := sdl.Point{X: 0, Y: -100}
	p3 := sdl.Point{X: -70, Y: -70}
	p5 := sdl.Point{X: -100, Y: 0}
	p6 := sdl.Point{X: -70, Y: 70}
	//    p0 := sdl.Point{X: -30, Y: 70}
	return []sdl.Point{p1, p, p2, p7, p4, p3, p5, p6}
}

func calcDist(points []sdl.Point, renderer *sdl.Renderer) {
	var lowest float64
	var lowest_idx int
	for idx, p1 := range points {
		lowest = 200
		lowest_idx = 0
	inner:
		for jd, p2 := range points {
			if idx == jd {
				continue inner
			}
			distance := distBetwPoints(p1, p2)
			if distance < lowest {
				lowest = distance
				lowest_idx = jd
			}
		}
		renderer.DrawLine(p1.X, p1.Y, points[lowest_idx].X, points[lowest_idx].Y)
	}
}

func subtractive(points []sdl.Point, renderer *sdl.Renderer) {
	for _, p1 := range points {
		for _, p2 := range points {
			if rand.Intn(100) > 60 {
				renderer.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
			}
		}
	}
}

func subtractiveMixed(points []sdl.Point, renderer *sdl.Renderer) {
	for _, p1 := range points {
		for _, p2 := range points {
			renderer.DrawLine(p1.X, p2.X, p1.Y, p2.Y)
		}
	}
}

func offsetPoints(points []sdl.Point, window *sdl.Window) []sdl.Point {
	W, H := window.GetSize()
	for idx := range points {
		points[idx].X += W / 4
		points[idx].Y += H / 4
	}
	return points
}

func offsetPointsOfLine(lines []Line) []Line {
	for idx := range lines {
		lines[idx].x1 += 400
		lines[idx].y1 += 400
		lines[idx].x2 += 400
		lines[idx].y2 += 400
	}
	return lines
}
