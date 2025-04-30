package main

import "github.com/hajimehoshi/ebiten/v2"

type Squirrel struct {
	X float64
	Y float64
}

func NewSquirrel() Squirrel {
	return Squirrel{
		X: 160,
		Y: 360, // Adjusted to appear walking on Mario-style grass
	}
}

func (g *Game) UpdateSquirrel() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.squirrel.X -= 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.squirrel.X += 4
	}
}
