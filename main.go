package main

import (
	"runtime"
	ct "tetris-go/constants"
	"tetris-go/tile"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	runtime.LockOSThread()

	rl.InitWindow(ct.SCREEN_WIDTH, ct.SCREEN_HEIGHT, "Tetris")
	rl.SetTargetFPS(60)

	tile1 := tile.NewTile(rl.Vector2{X: 10, Y: 10}, rl.Rectangle{Width: ct.TILE_SIZE, Height: ct.TILE_SIZE})

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(ct.Palette["Sky"])
		tile1.Draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
