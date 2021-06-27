package asset

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
)

//go:embed ball.png
var ballImgBytes []byte

func GetBallImg() (image.Image, error) {
	return png.Decode(bytes.NewBuffer(ballImgBytes))
}
