package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Config struct {
	GameName     string
	ScreenWidth  int
	ScreenHeight int
}

type game struct {
	gameName     string
	screenWidth  int
	screenHeight int
	playerPaddle *paddle
	botPaddle    *paddle
}

func New(cfg Config) ebiten.Game {
	g := &game{
		screenHeight: cfg.ScreenHeight,
		screenWidth:  cfg.ScreenWidth,
		gameName:     cfg.GameName,
	}
	g.init()
	return g
}

func (g *game) init() {
	ebiten.SetWindowSize(g.screenWidth, g.screenHeight)
	ebiten.SetWindowTitle("Tennis game")
	g.botPaddle = newPaddle(float64(g.screenWidth)/2-float64(paddleWidth)/2, 0)
	g.playerPaddle = newPaddle(float64(g.screenWidth)/2-float64(paddleWidth)/2, float64(g.screenHeight)-float64(paddleHeight))
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{})
	botPaddleOp := &ebiten.DrawImageOptions{}
	botPaddleOp.GeoM.Translate(g.botPaddle.x, g.botPaddle.y)
	screen.DrawImage(g.botPaddle.Image, botPaddleOp)

	playerPaddleOp := &ebiten.DrawImageOptions{}
	playerPaddleOp.GeoM.Translate(g.playerPaddle.x, g.playerPaddle.y)
	screen.DrawImage(g.playerPaddle.Image, playerPaddleOp)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
