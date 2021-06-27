package asset

import (
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/pkg/errors"
	gofont "golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	normalFontSize = 24
)

type Font interface {
	GetNormal() gofont.Face
}

type font struct {
	normal gofont.Face
}

func NewFont() (Font, error) {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't parse the font")
	}

	f := &font{}

	const dpi = 72
	f.normal, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    normalFontSize,
		DPI:     dpi,
		Hinting: gofont.HintingFull,
	})
	if err != nil {
		return nil, errors.Wrap(err, "couldn't parse create a face (normal)")
	}

	return f, nil
}

func (f *font) GetNormal() gofont.Face {
	return f.normal
}
