package main

import (
	"math"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Sniffle Shoots Asteroids", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		800, 800, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	// points := offsetPoints(createShape())
	points := offsetPoints(createTriangle())

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				sdl.Quit()
				return
				running = false
			}
		}
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.DrawPoint(400, 400)
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.DrawPoints(points)

		// calcDist(points, renderer)
		drawTriangle(points, renderer)
		renderer.Present()
		// subtractive(points, renderer)
		//	subtractiveMixed(points, renderer)
	}
}

func drawTriangle(points []sdl.Point, renderer *sdl.Renderer) {
}

func createTriangle() []sdl.Point {
	points := []sdl.Point{}
	p := sdl.Point{X: 100, Y: 0}
	p1 := sdl.Point{X: 0, Y: -100}
	p2 := sdl.Point{X: -100, Y: 0}
	points = append(points, p, p1, p2)
	return points
}

func createShape() []sdl.Point {
	points := []sdl.Point{}
	p := sdl.Point{X: 70, Y: 70}
	p1 := sdl.Point{X: 100, Y: 0}
	p2 := sdl.Point{X: 0, Y: 100}
	p3 := sdl.Point{X: -70, Y: -70}
	p4 := sdl.Point{X: 0, Y: -100}
	p5 := sdl.Point{X: -100, Y: 0}
	p6 := sdl.Point{X: -70, Y: 70}
	p7 := sdl.Point{X: 70, Y: -70}
	points = append(points, p, p1, p2, p3, p4, p5, p6, p7)
	return points
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
			distance := math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) + math.Pow(float64(p1.Y-p2.Y), 2))
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

func offsetPoints(points []sdl.Point) []sdl.Point {
	for idx := range points {
		points[idx].X += 400
		points[idx].Y += 400
	}
	return points
}
