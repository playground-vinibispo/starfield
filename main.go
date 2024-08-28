package main

import (
	"starfield/internals/helpers"
	"starfield/internals/models"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var stars [800]models.Star
	rl.InitWindow(700, 700, "Starfield")
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())
	rl.SetTargetFPS(60)

	for i := 0; i < len(stars); i++ {
		stars[i] = *models.NewStar()
	}

	for !rl.WindowShouldClose() {
		speed := helpers.MapValue(rl.GetMousePosition().X, 0, float32(rl.GetScreenWidth()), 0, 50)
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.BeginMode2D(rl.Camera2D{
			Target:   rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
			Offset:   rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
			Rotation: 0,
			Zoom:     1.0,
		})
		rl.BeginBlendMode(rl.BlendAlpha)
		for i := 0; i < len(stars); i++ {
			stars[i].Update(speed)
			stars[i].Show()
		}
		rl.EndBlendMode()
		rl.EndMode2D()
		rl.EndDrawing()
	}
}
