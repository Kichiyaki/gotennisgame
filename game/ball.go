package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pkg/errors"

	"github.com/Kichiyaki/gotennisgame/game/assets"
)

type ball struct {
	*ebiten.Image
	*coords
}

func newBall(x, y float64) (*ball, error) {
	ballImg, err := assets.GetBallImg()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't load ball img")
	}
	return &ball{
		Image: ebiten.NewImageFromImage(ballImg),
		coords: &coords{
			x: x,
			y: y,
		},
	}, nil
}
