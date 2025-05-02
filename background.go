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

	g.sunAngle += 0.02
	if g.sunAngle > 6.28 {
		g.sunAngle = 0
	}
	g.sunY = 20 + 5*math.Sin(g.sunAngle)
}

func (g *Game) DrawBackground(screen *ebiten.Image) {

	bgOpts := &ebiten.DrawImageOptions{}
	bgOpts.GeoM.Scale(
		640.0/float64(backgroundImg.Bounds().Dx()),
		480.0/float64(backgroundImg.Bounds().Dy()),
	)
	screen.DrawImage(backgroundImg, bgOpts)

	sunOpts := &ebiten.DrawImageOptions{}
	sunOpts.GeoM.Scale(0.18, 0.18)
	sunOpts.GeoM.Translate(g.sunX, g.sunY)
	screen.DrawImage(sunImg, sunOpts)

	cloud1 := &ebiten.DrawImageOptions{}
	cloud1.GeoM.Scale(0.4, 0.4)
	cloud1.GeoM.Translate(g.cloud1X, 50)
	screen.DrawImage(cloudImg, cloud1)

	cloud2 := &ebiten.DrawImageOptions{}
	cloud2.GeoM.Scale(0.4, 0.4)
	cloud2.GeoM.Translate(g.cloud2X, 100)
	screen.DrawImage(cloudImg, cloud2)

	treeOpts := &ebiten.DrawImageOptions{}
	treeOpts.GeoM.Scale(0.45, 0.45)
	treeOpts.GeoM.Translate(120, 30)
	screen.DrawImage(treeImg, treeOpts)

	grassOpts := &ebiten.DrawImageOptions{}
	const grassScale = 0.25
	grassOpts.GeoM.Scale(
		grassScale, grassScale,
	)
	grassOffsetY := float64(screen.Bounds().Dy()) - (grassScale * float64(grassImg.Bounds().Dy())) + 10
	grassOpts.GeoM.Translate(-10, grassOffsetY)
	screen.DrawImage(grassImg, grassOpts)
	grassOpts.GeoM.Translate(grassScale*float64(grassImg.Bounds().Dx())-40, 0)
	screen.DrawImage(grassImg, grassOpts)
}
