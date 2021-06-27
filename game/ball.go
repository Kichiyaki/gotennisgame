package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pkg/errors"

	"github.com/Kichiyaki/gotennisgame/game/asset"
)

const (
	ballMoveSpeed = 3
)

type ball struct {
	*renderableEntity
	speed *speed
}

func newBall(x, y float64) (*ball, error) {
	ballImg, err := asset.GetBallImg()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't load the ball img")
	}
	return &ball{
		renderableEntity: &renderableEntity{
			Image:        ebiten.NewImageFromImage(ballImg),
			coords:       newCoords(x, y),
			initPosition: newCoords(x, y),
		},
		speed: &speed{
			newCoords(ballMoveSpeed, ballMoveSpeed),
		},
	}, nil
}
