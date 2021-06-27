package game

import "github.com/hajimehoshi/ebiten/v2"

type renderable interface {
	getX() float64
	getY() float64
	getImage() *ebiten.Image
}

type renderableEntity struct {
	*ebiten.Image
	*coords
	initPosition *coords
}

func (r *renderableEntity) getImage() *ebiten.Image {
	return r.Image
}

func (r *renderableEntity) isPointInBoundaries(x, y float64) bool {
	return y > r.getY() && y < r.getBottomY() && x > r.getX() && x < r.getRightX()
}

func (r *renderableEntity) resetPosition() {
	r.setX(r.initPosition.getX())
	r.setY(r.initPosition.getY())
}

func (r *renderableEntity) getLeftX() float64 {
	return r.getX()
}

func (r *renderableEntity) getMidX() float64 {
	width, _ := r.Size()
	return r.getX() + float64(width)/2
}

func (r *renderableEntity) getRightX() float64 {
	width, _ := r.Size()
	return r.getX() + float64(width)
}

func (r *renderableEntity) getTopY() float64 {
	return r.getY()
}

func (r *renderableEntity) getBottomY() float64 {
	_, height := r.Size()
	return r.getY() + float64(height)
}
