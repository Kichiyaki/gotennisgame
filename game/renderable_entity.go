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
	entityWidth, entityHeight := r.Size()
	return y > r.getY() && y < r.getY()+float64(entityHeight) && x > r.getX() && x < r.getX()+float64(entityWidth)
}

func (r *renderableEntity) resetPosition() {
	r.setX(r.initPosition.getX())
	r.setY(r.initPosition.getY())
}
