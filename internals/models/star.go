package models

import (
	"starfield/internals/helpers"

	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type Star struct {
	x  float32
	y  float32
	z  float32
	pz float32
}

func random(min, max float32) float32 {
	return rand.Float32()*(max-min) + min
}

func NewStar() *Star {
	width := rl.GetScreenWidth()
	height := rl.GetScreenHeight()
	x := random(-float32(width), float32(width))
	y := random(-float32(height), float32(height))
	z := float32(width)
	pz := z
	return &Star{x: x, y: y, z: z, pz: pz}
}

func (s *Star) Update(speed float32) {
	width := rl.GetScreenWidth()
	height := rl.GetScreenHeight()
	s.z = s.z - speed
	if s.z < 1 {
		s.z = float32(width / 2)
		s.x = random(-float32(width/2), float32(width/2))
		s.y = random(-float32(height/2), float32(height/2))
		s.pz = s.z
	}
}

func (s *Star) Show() {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())
	sx := helpers.MapValue(s.x/s.z, 0, 1, 0, screenWidth)
	sy := helpers.MapValue(s.y/s.z, 0, 1, 0, screenHeight)
	r := helpers.MapValue(s.z, 0, screenWidth, 16, 0)
	px := helpers.MapValue(s.x/s.pz, 0, 1, 0, screenWidth)
	py := helpers.MapValue(s.y/s.pz, 0, 1, 0, screenHeight)
	rl.DrawCircleV(rl.NewVector2(sx, sy), r, rl.White)
	rl.DrawLineV(rl.NewVector2(px, py), rl.NewVector2(sx, sy), rl.Fade(rl.White, 0.5))
	s.pz = s.z
}
