package main

import (
	_ "image/png"
	"log"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	backgroundImg *ebiten.Image
	cloudImg      *ebiten.Image
	treeImg       *ebiten.Image
	squirrelImg   *ebiten.Image
	acornImg      *ebiten.Image
	sunImg        *ebiten.Image
)

type Game struct {
	squirrel     Squirrel
	acorns       []Acorn
	cloud1X      float64
	cloud2X      float64
	sunX         float64
	sunDirection float64
	score        int
	startTime    time.Time
	gameOver     bool
	timeElapsed  float64
}

func (g *Game) Update() error {
	if g.gameOver && ebiten.IsKeyPressed(ebiten.KeyR) {
		*g = Game{
			squirrel:     NewSquirrel(),
			acorns:       SpawnAcorns(),
			cloud1X:      0,
			cloud2X:      -300,
			sunX:         20,
			sunDirection: 1,
			score:        0,
			startTime:    time.Now(),
		}
		return nil
	}

	if g.gameOver {
		return nil
	}

	g.UpdateSquirrel()
	g.UpdateAcorns()

	g.cloud1X += 0.3
	g.cloud2X += 0.4
	if g.cloud1X > 700 {
		g.cloud1X = -100
	}
	if g.cloud2X > 700 {
		g.cloud2X = -100
	}

	g.sunX += g.sunDirection * 0.2
	if g.sunX > 80 {
		g.sunDirection = -1
	} else if g.sunX < 20 {
		g.sunDirection = 1
	}

	g.timeElapsed = time.Since(g.startTime).Seconds()
	if g.timeElapsed >= 90 {
		g.gameOver = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	bgOpts := &ebiten.DrawImageOptions{}
	bgOpts.GeoM.Scale(
		640.0/float64(backgroundImg.Bounds().Dx()),
		480.0/float64(backgroundImg.Bounds().Dy()),
	)
	screen.DrawImage(backgroundImg, bgOpts)

	sunOpts := &ebiten.DrawImageOptions{}
	sunOpts.GeoM.Scale(0.2, 0.2)
	sunOpts.GeoM.Translate(g.sunX, 20)
	screen.DrawImage(sunImg, sunOpts)

	cloudOpts1 := &ebiten.DrawImageOptions{}
	cloudOpts1.GeoM.Translate(g.cloud1X, 50)
	screen.DrawImage(cloudImg, cloudOpts1)

	cloudOpts2 := &ebiten.DrawImageOptions{}
	cloudOpts2.GeoM.Translate(g.cloud2X, 100)
	screen.DrawImage(cloudImg, cloudOpts2)

	treeOpts := &ebiten.DrawImageOptions{}
	treeOpts.GeoM.Scale(0.4, 0.4)
	treeOpts.GeoM.Translate(100, 50)
	screen.DrawImage(treeImg, treeOpts)

	squirrelOpts := &ebiten.DrawImageOptions{}
	squirrelOpts.GeoM.Scale(0.2, 0.2)
	squirrelOpts.GeoM.Translate(g.squirrel.X, g.squirrel.Y)
	screen.DrawImage(squirrelImg, squirrelOpts)

	g.DrawAcorns(screen)

	if g.gameOver {
		ebitenutil.DebugPrint(screen, "GAME OVER\nFinal Score: "+strconv.Itoa(g.score)+"\nPress R to Restart")
	} else {
		remaining := int(90 - g.timeElapsed)
		ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(g.score)+"   Time left: "+strconv.Itoa(remaining)+"s")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal("Error loading image:", path, err)
	}
	return ebiten.NewImageFromImage(img)
}

func main() {
	backgroundImg = loadImage("assets/background.png")
	cloudImg = loadImage("assets/cloud.png")
	treeImg = loadImage("assets/tree.png")
	squirrelImg = loadImage("assets/squirrel.png")
	acornImg = loadImage("assets/acorn.png")
	sunImg = loadImage("assets/sun.png")

	if err := ebiten.RunGame(&Game{
		squirrel:     NewSquirrel(),
		acorns:       SpawnAcorns(),
		cloud1X:      0,
		cloud2X:      -300,
		sunX:         20,
		sunDirection: 1,
		score:        0,
		startTime:    time.Now(),
	}); err != nil {
		log.Fatal(err)
	}
}
