package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// UpdateBackground handles movement of the clouds and sun (bobbing effect).
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

	// Bob the sun in place like Paper Mario
	g.sunAngle += 0.02
	if g.sunAngle > 6.28 {
		g.sunAngle = 0
	}
	g.sunY = 20 + 5*math.Sin(g.sunAngle)
}

// DrawBackground draws the background image, sun, clouds, and tree.
func (g *Game) DrawBackground(screen *ebiten.Image) {
	// Draw background image
	bgOpts := &ebiten.DrawImageOptions{}
	bgOpts.GeoM.Scale(
		640.0/float64(backgroundImg.Bounds().Dx()),
		480.0/float64(backgroundImg.Bounds().Dy()),
	)
	screen.DrawImage(backgroundImg, bgOpts)

	// Draw sun
	sunOpts := &ebiten.DrawImageOptions{}
	sunOpts.GeoM.Scale(0.2, 0.2)
	sunOpts.GeoM.Translate(g.sunX, g.sunY)
	screen.DrawImage(sunImg, sunOpts)

	// Draw clouds
	cloudOpts1 := &ebiten.DrawImageOptions{}
	cloudOpts1.GeoM.Translate(g.cloud1X, 50)
	screen.DrawImage(cloudImg, cloudOpts1)

	cloudOpts2 := &ebiten.DrawImageOptions{}
	cloudOpts2.GeoM.Translate(g.cloud2X, 100)
	screen.DrawImage(cloudImg, cloudOpts2)

	// Draw tree
	treeOpts := &ebiten.DrawImageOptions{}
	treeOpts.GeoM.Scale(0.4, 0.4)
	treeOpts.GeoM.Translate(100, 50)
	screen.DrawImage(treeImg, treeOpts)
}
