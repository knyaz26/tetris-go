package main

import (
	"runtime"

	"tetris-go/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	runtime.LockOSThread()

	rl.InitWindow(constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT, "Tetris")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
