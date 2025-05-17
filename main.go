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

	//this returns a slice of type tiles.
	tiles := drawChamber()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(ct.Palette["Sky"])

		drawTiles(tiles)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

//***************************************************************************************************//

func drawChamber() []*tile.Tile {
	tiles := []*tile.Tile{}
	for i := 0; i < ct.CHAMBER_HEIGHT; i++ {
		for j := 0; j < ct.CHAMBER_WIDTH; j++ {
			tilePos := rl.Vector2{
				X: float32(ct.CHAMBER_POS_X + (j * ct.TILE_SIZE)),
				Y: float32(ct.CHAMBER_POS_Y + (i * ct.TILE_SIZE)),
			}
			tileSize := rl.Rectangle{
				X:      tilePos.X,
				Y:      tilePos.Y,
				Width:  float32(ct.TILE_SIZE),
				Height: float32(ct.TILE_SIZE),
			}
			newTile := tile.NewTile(tilePos, tileSize)
			tiles = append(tiles, newTile)
		}
	}
	return tiles
}

func drawTiles(tiles []*tile.Tile) {
	for _, t := range tiles {
		t.Draw()
	}
}
