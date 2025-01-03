package game

import (
	"RPG-Brazil/graphics"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Game struct {
	renderer *graphics.Renderer
	gameMap  *GameMap
}

func NewGame(renderer *graphics.Renderer) *Game {
	gameMap := NewGameMap(10, 10) // Example map size
	return &Game{renderer: renderer, gameMap: gameMap}
}

func (g *Game) Update() {
	// Handle input and update game state
	g.handleInput()
}

func (g *Game) Render() {
	g.gameMap.Render(g.renderer)
}

func (g *Game) handleInput() {
	// Handle WASD and mouse input for movement
	// This is a placeholder for actual input handling logic
	if glfw.Press == glfw.GetKey(glfw.KeyW) {
		// Move forward
	}
	if glfw.Press == glfw.GetKey(glfw.KeyA) {
		// Move left
	}
	if glfw.Press == glfw.GetKey(glfw.KeyS) {
		// Move backward
	}
	if glfw.Press == glfw.GetKey(glfw.KeyD) {
		// Move right
	}
	// Handle mouse input for looking around
}
