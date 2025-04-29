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
	sunY         float64
	sunAngle     float64
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
			sunY:         20,
			sunAngle:     0,
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
	g.UpdateBackground()

	g.timeElapsed = time.Since(g.startTime).Seconds()
	if g.timeElapsed >= 90 {
		g.gameOver = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawBackground(screen)

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
		sunY:         20,
		sunAngle:     0,
		sunDirection: 1,
		score:        0,
		startTime:    time.Now(),
	}); err != nil {
		log.Fatal(err)
	}
}
