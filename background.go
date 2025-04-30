package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) UpdateBackground() {
	// Move clouds
	g.cloud1X += 0.3
	g.cloud2X += 0.4
	if g.cloud1X > 700 {
		g.cloud1X = -100
	}
	if g.cloud2X > 700 {
		g.cloud2X = -100
	}

	// Sun bobbing
	g.sunAngle += 0.02
	if g.sunAngle > 6.28 {
		g.sunAngle = 0
	}
	g.sunY = 20 + 5*math.Sin(g.sunAngle)
}

func (g *Game) DrawBackground(screen *ebiten.Image) {
	// Sky background
	bgOpts := &ebiten.DrawImageOptions{}
	bgOpts.GeoM.Scale(
		640.0/float64(backgroundImg.Bounds().Dx()),
		480.0/float64(backgroundImg.Bounds().Dy()),
	)
	screen.DrawImage(backgroundImg, bgOpts)

	// Sun
	sunOpts := &ebiten.DrawImageOptions{}
	sunOpts.GeoM.Scale(0.18, 0.18)
	sunOpts.GeoM.Translate(g.sunX, g.sunY)
	screen.DrawImage(sunImg, sunOpts)

	// Clouds
	cloud1 := &ebiten.DrawImageOptions{}
	cloud1.GeoM.Scale(0.4, 0.4)
	cloud1.GeoM.Translate(g.cloud1X, 50)
	screen.DrawImage(cloudImg, cloud1)

	cloud2 := &ebiten.DrawImageOptions{}
	cloud2.GeoM.Scale(0.4, 0.4)
	cloud2.GeoM.Translate(g.cloud2X, 100)
	screen.DrawImage(cloudImg, cloud2)

	// Tree
	treeOpts := &ebiten.DrawImageOptions{}
	treeOpts.GeoM.Scale(0.45, 0.45)
	treeOpts.GeoM.Translate(120, 100)
	screen.DrawImage(treeImg, treeOpts)

	// ✅ Grass — Mario-style, shows dirt, raised up a little
	grassOpts := &ebiten.DrawImageOptions{}
	grassOpts.GeoM.Scale(
		640.0/float64(grassImg.Bounds().Dx()),
		140.0/float64(grassImg.Bounds().Dy()),
	)
	grassOpts.GeoM.Translate(0, 360) // Move up to show base
	screen.DrawImage(grassImg, grassOpts)
}
