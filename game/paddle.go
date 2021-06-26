package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	paddleWidth  = 100
	paddleHeight = 20
)

var paddleColor = color.RGBA{
	R: 255,
	G: 255,
	B: 255,
	A: 255,
}

type paddle struct {
	*ebiten.Image
	*coords
}

func newPaddle(x, y float64) *paddle {
	return &paddle{
		Image: newRectangle(paddleWidth, paddleHeight, paddleColor),
		coords: &coords{
			x: x,
			y: y,
		},
	}
}

func (p *paddle) isPointInBoundaries(x, y float64) bool {
	return y > p.getY() && y < p.getY()+paddleHeight && x > p.getX() && x < p.getX()+paddleWidth
}
