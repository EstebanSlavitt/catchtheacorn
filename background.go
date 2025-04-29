package main

import "github.com/hajimehoshi/ebiten/v2"

// UpdateBackground moves the clouds and sun smoothly across the screen.
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

	// Move sun side-to-side
	g.sunX += g.sunDirection * 0.2
	if g.sunX > 80 {
		g.sunDirection = -1
	} else if g.sunX < 20 {
		g.sunDirection = 1
	}
}

// DrawBackground draws the static and animated background elements.
func (g *Game) DrawBackground(screen *ebiten.Image) {
	// Draw background sky image
	bgOpts := &ebiten.DrawImageOptions{}
	bgOpts.GeoM.Scale(
		640.0/float64(backgroundImg.Bounds().Dx()),
		480.0/float64(backgroundImg.Bounds().Dy()),
	)
	screen.DrawImage(backgroundImg, bgOpts)

	// Draw sun
	sunOpts := &ebiten.DrawImageOptions{}
	sunOpts.GeoM.Scale(0.2, 0.2)
	sunOpts.GeoM.Translate(g.sunX, 20)
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
