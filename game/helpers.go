package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func newRectangle(width, height int, c color.Color) *ebiten.Image {
	img := ebiten.NewImage(paddleWidth, paddleHeight)
	img.Fill(paddleColor)
	return img
}
