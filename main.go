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
	x      float64
	y      float64
	isMega bool
}

var (
	backgroundImg *ebiten.Image
	cloudImg      *ebiten.Image
	treeImg       *ebiten.Image
	squirrelImg   *ebiten.Image
	acornImg      *ebiten.Image
	sunImg        *ebiten.Image
)

type Game struct {
	playerX      float64
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
			playerX:      160,
			acorns:       spawnAcorns(),
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

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerX -= 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerX += 4
	}

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

	for i := range g.acorns {
		speed := 1.5
		if g.acorns[i].isMega {
			speed = 2.2
		}
		g.acorns[i].y += speed

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
			if g.acorns[i].isMega {
				g.score += 5
			} else {
				g.score++
			}
			g.acorns[i] = g.newAcorn()
		}

		if g.acorns[i].y > 480 {
			g.acorns[i] = g.newAcorn()
		}
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

	
	treeOpts := &ebiten.DrawImageOptions{}
	treeOpts.GeoM.Scale(0.4, 0.4)
	treeOpts.GeoM.Translate(100, 50)
	screen.DrawImage(treeImg, treeOpts)

	
	squirrelOpts := &ebiten.DrawImageOptions{}
	squirrelOpts.GeoM.Scale(0.2, 0.2)
	squirrelOpts.GeoM.Translate(g.playerX, 340)
	screen.DrawImage(squirrelImg, squirrelOpts)

	 
	cloudOpts1 := &ebiten.DrawImageOptions{}
	cloudOpts1.GeoM.Translate(g.cloud1X, 50)
	screen.DrawImage(cloudImg, cloudOpts1)

	cloudOpts2 := &ebiten.DrawImageOptions{}
	cloudOpts2.GeoM.Translate(g.cloud2X, 100)
	screen.DrawImage(cloudImg, cloudOpts2)

	
	sunOpts := &ebiten.DrawImageOptions{}
	sunOpts.GeoM.Scale(0.2, 0.2)
	sunOpts.GeoM.Translate(g.sunX, 20)
	screen.DrawImage(sunImg, sunOpts)

	// Acorns (drawn LAST so they are visible!)
	for _, acorn := range g.acorns {
		acornOpts := &ebiten.DrawImageOptions{}
		scale := 0.07
		if acorn.isMega {
			scale = 0.1
		}
		acornOpts.GeoM.Scale(scale, scale)
		acornOpts.GeoM.Translate(acorn.x, acorn.y)
		screen.DrawImage(acornImg, acornOpts)
	}

	
	if g.gameOver {
		ebitenutil.DebugPrint(screen, "GAME OVER\nFinal Score: "+strconv.Itoa(g.score)+"\nPress R to Restart")
	} else {
		remaining := int(90 - g.timeElapsed)
		ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(g.score)+"   Time left: "+strconv.Itoa(remaining)+"s")
	}
}

func (g *Game) newAcorn() Acorn {
	return Acorn{
		x:      float64(rand.Intn(600)),
		y:      0,
		isMega: rand.Float64() < 0.1,
	}
}

func spawnAcorns() []Acorn {
	return []Acorn{
		{x: float64(rand.Intn(600)), y: 0, isMega: rand.Float64() < 0.1},
		{x: float64(rand.Intn(600)), y: -150, isMega: rand.Float64() < 0.1},
		{x: float64(rand.Intn(600)), y: -300, isMega: rand.Float64() < 0.1},
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

	backgroundImg = loadImage("assets/background.png")
	cloudImg = loadImage("assets/cloud.png")
	treeImg = loadImage("assets/tree.png")
	squirrelImg = loadImage("assets/squirrel.png")
	acornImg = loadImage("assets/acorn.png")
	sunImg = loadImage("assets/sun.png")

	if err := ebiten.RunGame(&Game{
		playerX:      160,
		acorns:       spawnAcorns(),
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
