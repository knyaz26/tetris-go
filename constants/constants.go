package constants

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	TILE_SIZE      = 32
	SCREEN_WIDTH   = 640
	SCREEN_HEIGHT  = 700
	CHAMBER_WIDTH  = 10
	CHAMBER_HEIGHT = 20
)

var (
	//this pos can be used to center the chamber in the screen.
	//CHAMBER_POS_X = (SCREEN_WIDTH - (CHAMBER_WIDTH * TILE_SIZE)) / 2
	CHAMBER_POS_X = SCREEN_WIDTH / 20
	CHAMBER_POS_Y = (SCREEN_HEIGHT - (CHAMBER_HEIGHT * TILE_SIZE)) / 2
)

var Palette = map[string]rl.Color{
	"Red":    rl.NewColor(235, 64, 52, 255),
	"Green":  rl.NewColor(113, 201, 50, 255),
	"Blue":   rl.NewColor(43, 26, 107, 255),
	"Sky":    rl.NewColor(50, 131, 201, 255),
	"Yellow": rl.NewColor(218, 245, 44, 255),
	"Orange": rl.NewColor(247, 141, 2, 255),
	"Gray":   rl.NewColor(50, 50, 50, 255),
	"White":  rl.NewColor(255, 255, 255, 255),
	"Black":  rl.NewColor(0, 0, 0, 255),
}
