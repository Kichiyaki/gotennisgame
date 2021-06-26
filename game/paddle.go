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
	x float64
	y float64
}

func newPaddle(x, y float64) *paddle {
	return &paddle{
		Image: newRectangle(paddleWidth, paddleHeight, paddleColor),
		x:     x,
		y:     y,
	}
}
