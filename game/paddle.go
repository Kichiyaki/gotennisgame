package game

import (
	"image/color"
)

const (
	paddleWidth             = 100
	paddleHeight            = 20
	paddleComputerMoveSpeed = 3
)

var paddleColor = color.White

type paddle struct {
	*renderableEntity
	speed *speed
}

func newPaddle(x, y float64) *paddle {
	return &paddle{
		renderableEntity: &renderableEntity{
			Image:        newRectangle(paddleWidth, paddleHeight, paddleColor),
			coords:       newCoords(x, y),
			initPosition: newCoords(x, y),
		},
		speed: &speed{
			coords: newCoords(paddleComputerMoveSpeed, paddleComputerMoveSpeed),
		},
	}
}
