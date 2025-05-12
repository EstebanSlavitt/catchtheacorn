package main

import (
	"math/rand"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

type Acorn struct {
	X      float64
	Y      float64
	IsMega bool
	IsBomb bool
}

func NewAcorn() *Acorn {
	r := rand.Float64()
	return &Acorn{
		X:      float64(rand.Intn(600)),
		Y:      0,
		IsMega: r < 0.1,
		IsBomb: r >= 0.9,
	}
}

func SpawnAcorns() []*Acorn {
	count := rand.Intn(4) + 2
	acorns := make([]*Acorn, count)
	for i := range acorns {
		acorns[i] = NewAcorn()
	}
	return acorns
}

func (g *Game) UpdateAcorns() {
	for i := range g.acorns {
		speed := 1.5
		if g.acorns[i].IsMega {
			speed = 2.2
		}
		g.acorns[i].Y += speed

		squirrelWidth := 80.0
		squirrelHeight := 80.0
		acornWidth := 30.0
		acornHeight := 30.0

		squirrelRect := struct {
			x, y, w, h float64
		}{
			x: g.squirrel.X,
			y: g.squirrel.Y,
			w: squirrelWidth,
			h: squirrelHeight,
		}

		acornRect := struct {
			x, y, w, h float64
		}{
			x: g.acorns[i].X,
			y: g.acorns[i].Y,
			w: acornWidth,
			h: acornHeight,
		}

		if squirrelRect.x < acornRect.x+acornRect.w &&
			squirrelRect.x+squirrelRect.w > acornRect.x &&
			squirrelRect.y < acornRect.y+acornRect.h &&
			squirrelRect.y+squirrelRect.h > acornRect.y {

			if g.acorns[i].IsBomb {
				if g.score >= 5 {
					g.score -= 5
				} else {
					g.score = 0
				}
			} else if g.acorns[i].IsMega {
				g.score += 5
			} else {
				g.score++
			}
			g.acorns[i] = nil
		}

		if g.acorns[i] != nil && g.acorns[i].Y > 400 {
			g.acorns[i] = nil
		}
	}

	// removing empty elements from the array
	g.acorns = slices.DeleteFunc(g.acorns, func(acorn *Acorn) bool {
		return acorn == nil
	})

	if rand.Float64() < 0.01 && len(g.acorns) < 8 {
		g.acorns = append(g.acorns, NewAcorn())
	}
}

func (g *Game) DrawAcorns(screen *ebiten.Image) {
	for _, acorn := range g.acorns {
		acornOpts := &ebiten.DrawImageOptions{}

		scale := 0.05
		if acorn.IsMega {
			scale = 0.1
		}
		acornOpts.GeoM.Scale(scale, scale)
		acornOpts.GeoM.Translate(acorn.X, acorn.Y)

		var img *ebiten.Image
		if acorn.IsBomb {
			img = bombImg
		} else if acorn.IsMega {
			img = megaAcornImg
		} else {
			img = acornImg
		}

		screen.DrawImage(img, acornOpts)
	}
}
