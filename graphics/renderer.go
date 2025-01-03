package graphics

import (
	"RPG-Brazil/game"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Renderer struct {
	// Add renderer state variables here
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Render() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	// Add rendering logic here
}

func (r *Renderer) RenderTile(tile game.Tile) {
	// Placeholder for actual rendering logic
	// Use tile.Position and tile.Texture to render the tile
}
