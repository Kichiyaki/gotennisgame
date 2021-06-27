package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func newRectangle(width, height int, c color.Color) *ebiten.Image {
	img := ebiten.NewImage(width, height)
	img.Fill(c)
	return img
}
