package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Acorn struct {
	x float64
	y float64
}

var (
	treeImg     *ebiten.Image
	squirrelImg *ebiten.Image
	acornImg    *ebiten.Image
)

type Game struct {
	playerX     float64
	acorns      []Acorn
	score       int
	startTime   time.Time
	gameOver    bool
	timeElapsed float64
}

func (g *Game) Update() error {
	if g.gameOver {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerX -= 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerX += 4
	}

	// Update all acorns
	for i := range g.acorns {
		g.acorns[i].y += 1.5

		// Check for catch
		squirrelWidth := 80.0
		squirrelHeight := 80.0
		acornWidth := 30.0
		acornHeight := 30.0

		squirrelRect := struct {
			x, y, w, h float64
		}{
			x: g.playerX,
			y: 340,
			w: squirrelWidth,
			h: squirrelHeight,
		}

		acornRect := struct {
			x, y, w, h float64
		}{
			x: g.acorns[i].x,
			y: g.acorns[i].y,
			w: acornWidth,
			h: acornHeight,
		}

		if squirrelRect.x < acornRect.x+acornRect.w &&
			squirrelRect.x+squirrelRect.w > acornRect.x &&
			squirrelRect.y < acornRect.y+acornRect.h &&
			squirrelRect.y+squirrelRect.h > acornRect.y {
			g.score++
			g.acorns[i] = g.newAcorn()
		}

		// Reset acorn if missed (but don't reset score)
		if g.acorns[i].y > 480 {
			g.acorns[i] = g.newAcorn()
		}
	}

	// Timer check
	g.timeElapsed = time.Since(g.startTime).Seconds()
	if g.timeElapsed >= 90 {
		g.gameOver = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Tree
	treeOpts := &ebiten.DrawImageOptions{}
	treeOpts.GeoM.Scale(0.4, 0.4)
	treeOpts.GeoM.Translate(100, 50)
	screen.DrawImage(treeImg, treeOpts)

	// Squirrel
	squirrelOpts := &ebiten.DrawImageOptions{}
	squirrelOpts.GeoM.Scale(0.2, 0.2)
	squirrelOpts.GeoM.Translate(g.playerX, 340)
	screen.DrawImage(squirrelImg, squirrelOpts)

	// Acorns
	for _, acorn := range g.acorns {
		acornOpts := &ebiten.DrawImageOptions{}
		acornOpts.GeoM.Scale(0.07, 0.07)
		acornOpts.GeoM.Translate(acorn.x, acorn.y)
		screen.DrawImage(acornImg, acornOpts)
	}

	// Score + Timer
	if g.gameOver {
		ebitenutil.DebugPrint(screen, "GAME OVER\nFinal Score: "+strconv.Itoa(g.score))
	} else {
		remaining := int(90 - g.timeElapsed)
		ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(g.score)+"   Time left: "+strconv.Itoa(remaining)+"s")
	}
}

func (g *Game) newAcorn() Acorn {
	return Acorn{
		x: float64(rand.Intn(600)),
		y: 0,
	}
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

	acorns := []Acorn{
		{x: float64(rand.Intn(600)), y: 0},
		{x: float64(rand.Intn(600)), y: -150},
		{x: float64(rand.Intn(600)), y: -300},
	}

	if err := ebiten.RunGame(&Game{
		playerX:   160,
		acorns:    acorns,
		score:     0,
		startTime: time.Now(),
	}); err != nil {
		log.Fatal(err)
	}
}
