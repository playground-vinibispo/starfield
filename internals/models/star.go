package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type Star struct {
	X  float32
	Y  float32
	Z  float32
	Pz float32
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
	return &Star{X: x, Y: y, Z: z, Pz: pz}
}

func (s *Star) Update(speed float32) {
	width := rl.GetScreenWidth()
	height := rl.GetScreenHeight()
	s.Z = s.Z - speed
	if s.Z < 1 {
		s.Z = float32(width)
		s.X = random(-float32(width), float32(width))
		s.Y = random(-float32(height), float32(height))
		s.Pz = s.Z
	}
}

func (s *Star) Show() {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())
	sx := rl.Remap(s.X/s.Z, 0, 1, 0, screenWidth)
	sy := rl.Remap(s.Y/s.Z, 0, 1, 0, screenHeight)
	px := rl.Remap(s.X/s.Pz, 0, 1, 0, screenWidth)
	py := rl.Remap(s.Y/s.Pz, 0, 1, 0, screenHeight)
	rl.DrawPixelV(rl.NewVector2(sx, sy), rl.White)
	rl.DrawLineV(rl.NewVector2(px, py), rl.NewVector2(sx, sy), rl.Fade(rl.White, 0.5))
	s.Pz = s.Z
}
