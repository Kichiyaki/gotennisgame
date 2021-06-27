package game

import (
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
	*renderableEntity
}

func newPaddle(x, y float64) *paddle {
	return &paddle{
		&renderableEntity{
			Image:        newRectangle(paddleWidth, paddleHeight, paddleColor),
			coords:       newCoords(x, y),
			initPosition: newCoords(x, y),
		},
	}
}
