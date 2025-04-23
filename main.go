package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png" // Enables PNG decoding
	"log"
	"os"
)

var (
	treeImg     *ebiten.Image
	squirrelImg *ebiten.Image
)

type Game struct {
	playerX float64
}

func (g *Game) Update() error {
	
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerX -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerX += 2
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	
	treeOpts := &ebiten.DrawImageOptions{}
	treeOpts.GeoM.Scale(0.3, 0.3)         
	treeOpts.GeoM.Translate(100, 50)      
	screen.DrawImage(treeImg, treeOpts)

	
	squirrelOpts := &ebiten.DrawImageOptions{}
	squirrelOpts.GeoM.Scale(0.2, 0.2)           
	squirrelOpts.GeoM.Translate(g.playerX, 340) 
	screen.DrawImage(squirrelImg, squirrelOpts)
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
	treeImg = loadImage("assets/tree.png")
	squirrelImg = loadImage("assets/squirrel.png")

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Catch the Acorn")

	if err := ebiten.RunGame(&Game{playerX: 160}); err != nil { 
		log.Fatal(err)
	}
}
