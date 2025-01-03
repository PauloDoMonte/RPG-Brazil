package game

import (
	"RPG-Brazil/graphics"

	"github.com/go-gl/mathgl/mgl32"
)

type Tile struct {
	Position mgl32.Vec3
	Texture  uint32
}

type GameMap struct {
	Tiles [][]Tile
}

func NewGameMap(width, height int) *GameMap {
	tiles := make([][]Tile, width)
	for i := range tiles {
		tiles[i] = make([]Tile, height)
		for j := range tiles[i] {
			tiles[i][j] = Tile{
				Position: mgl32.Vec3{float32(i), 0, float32(j)},
				Texture:  0, // Placeholder for texture ID
			}
		}
	}
	return &GameMap{Tiles: tiles}
}

func (gm *GameMap) Render(renderer *graphics.Renderer) {
	for _, row := range gm.Tiles {
		for _, tile := range row {
			renderer.RenderTile(tile)
		}
	}
}
