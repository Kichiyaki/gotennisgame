package game

import "sync"

type coords struct {
	mu sync.Mutex
	x  float64
	y  float64
}

func (c *coords) getX() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.x
}

func (c *coords) setX(x float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.x = x
}

func (c *coords) getY() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.y
}

func (c *coords) setY(y float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.y = y
}
