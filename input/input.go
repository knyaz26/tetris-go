package input

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func IsActionJustPressed(actionName string) bool {
	switch actionName {
	case "ui_left":
		return rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressed(rl.KeyA)
	case "ui_right":
		return rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressed(rl.KeyD)
	case "ui_down":
		return rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressed(rl.KeyS)
	case "ui_up":
		return rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyW)
	case "ui_accept":
		return rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter)
	default:
		return false
	}
}

func IsActionPressed(actionName string) bool {
	switch actionName {
	case "ui_left":
		return rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA)
	case "ui_right":
		return rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD)
	case "ui_down":
		return rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS)
	case "ui_up":
		return rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW)
	case "ui_accept":
		return rl.IsKeyDown(rl.KeySpace) || rl.IsKeyDown(rl.KeyEnter)
	default:
		return false
	}
}
