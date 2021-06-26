package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"sync"
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
	mu sync.Mutex
	x  float64
	y  float64
}

func newPaddle(x, y float64) *paddle {
	return &paddle{
		Image: newRectangle(paddleWidth, paddleHeight, paddleColor),
		x:     x,
		y:     y,
	}
}

func (p *paddle) getX() float64 {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.x
}

func (p *paddle) setX(x float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.x = x
}

func (p *paddle) getY() float64 {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.y
}

func (p *paddle) setY(y float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.y = y
}
