package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Star struct {
	x float32
	y float32
	z float32
}

func NewStar() *Star {
	width := rl.GetScreenWidth()
	height := rl.GetScreenHeight()
	x := rl.GetRandomValue(-int32(width), int32(width))
	y := rl.GetRandomValue(-int32(height), int32(height))
	z := width
	return &Star{float32(x), float32(y), float32(z)}
}

func (s *Star) Update(speed float32) {
	width := rl.GetScreenWidth()
	height := rl.GetScreenHeight()
	s.z = s.z - speed
	if s.z < 1 {
		s.z = float32(width)
		s.x = float32(rl.GetRandomValue(-int32(width), int32(width)))
		s.y = float32(rl.GetRandomValue(-int32(height), int32(height)))
	}
}

func mapValue(value, start1, stop1, start2, stop2 float32) float32 {
	newVal := (value-start1)/(stop1-start1)*(stop2-start2) + start2
	if start2 < stop2 {
		return constrain(newVal, start2, stop2)
	} else {
		return constrain(newVal, stop2, start2)
	}
}

func constrain(value, min, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func (s *Star) Show() {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())
	sx := mapValue(s.x/s.z, 0, 1, 0, screenWidth)
	sy := mapValue(s.y/s.z, 0, 1, 0, screenHeight)
	r := mapValue(s.z, 0, screenWidth, 16, 0)
	rl.DrawCircleV(rl.NewVector2(sx, sy), r, rl.White)
}

var speed float32

func main() {
	var stars [800]Star
	rl.InitWindow(700, 700, "Starfield")
	rl.SetTargetFPS(60)

	for i := 0; i < len(stars); i++ {
		stars[i] = *NewStar()
	}

	for !rl.WindowShouldClose() {
		speed := mapValue(rl.GetMousePosition().X, 0, float32(rl.GetScreenWidth()), 0, 20)
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for i := 0; i < len(stars); i++ {
			stars[i].Update(speed)
			stars[i].Show()
		}

		rl.EndDrawing()
	}
}
