package game

import (
	"image/color"
)

const (
	paddleComputerMoveSpeed = 3
)

var paddleColor = color.White

type paddle struct {
	*renderableEntity
	speed *speed
}

func newPaddle(x, y float64, width, height int) *paddle {
	return &paddle{
		renderableEntity: &renderableEntity{
			Image:        newRectangle(width, height, paddleColor),
			coords:       newCoords(x, y),
			initPosition: newCoords(x, y),
		},
		speed: &speed{
			coords: newCoords(paddleComputerMoveSpeed, paddleComputerMoveSpeed),
		},
	}
}
