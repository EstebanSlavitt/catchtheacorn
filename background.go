package main

import (
	"image/color"
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

	screen.Fill(color.RGBA{R: 109, G: 211, B: 217, A: 100})

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

	grassWidth := float64(grassImg.Bounds().Dx())
	grassY := float64(screen.Bounds().Dy()) - float64(grassImg.Bounds().Dy())

	for x := 0.0; x < 640; x += grassWidth {
		grassOpts := &ebiten.DrawImageOptions{}
		grassOpts.GeoM.Translate(x, grassY)
		screen.DrawImage(grassImg, grassOpts)
	}

}
