package tile

import (
	ct "tetris-go/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	Position rl.Vector2
	Size     rl.Rectangle
	Occupied bool
	Color    rl.Color
}

func NewTile(position rl.Vector2, size rl.Rectangle) *Tile {
	return &Tile{
		Position: position,
		Size:     size,
		Occupied: false,
		Color:    ct.Palette["Black"],
	}
}

func (t *Tile) Draw() {
	rl.DrawRectangleRec(rl.Rectangle{
		X:      t.Position.X,
		Y:      t.Position.Y,
		Width:  t.Size.Width,
		Height: t.Size.Height,
	}, t.Color)

	rl.DrawRectangleRec(rl.Rectangle{
		X:      t.Position.X + 2,
		Y:      t.Position.Y + 2,
		Width:  t.Size.Width - 4,
		Height: t.Size.Height - 4,
	}, ct.Palette["Gray"])
}
