package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	treeImg     *ebiten.Image
	squirrelImg *ebiten.Image
	acornImg    *ebiten.Image
)

type Game struct {
	playerX float64
	acornX  float64
	acornY  float64
}

func (g *Game) Update() error {
	// Faster squirrel movement
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerX -= 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerX += 4
	}

	// Slower acorn falling speed
	g.acornY += 1.5

	// Reset acorn if it hits the ground
	if g.acornY > 480 {
		g.acornY = 0
		g.acornX = float64(rand.Intn(600))
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Tree
	treeOpts := &ebiten.DrawImageOptions{}
	treeOpts.GeoM.Scale(0.3, 0.3)
	treeOpts.GeoM.Translate(100, 50)
	screen.DrawImage(treeImg, treeOpts)

	// Squirrel
	squirrelOpts := &ebiten.DrawImageOptions{}
	squirrelOpts.GeoM.Scale(0.2, 0.2)
	squirrelOpts.GeoM.Translate(g.playerX, 340)
	screen.DrawImage(squirrelImg, squirrelOpts)

	// Acorn
	acornOpts := &ebiten.DrawImageOptions{}
	acornOpts.GeoM.Scale(0.1, 0.1)
	acornOpts.GeoM.Translate(g.acornX, g.acornY)
	screen.DrawImage(acornImg, acornOpts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func loadImage(path string) *ebiten.Image {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening image:", path, err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal("Error decoding image:", path, err)
	}
	return ebiten.NewImageFromImage(img)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	treeImg = loadImage("assets/tree.png")
	squirrelImg = loadImage("assets/squirrel.png")
	acornImg = loadImage("assets/acorn.png")

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Catch the Acorn")

	if err := ebiten.RunGame(&Game{
		playerX: 160,
		acornX:  float64(rand.Intn(600)),
		acornY:  0,
	}); err != nil {
		log.Fatal(err)
	}
}
