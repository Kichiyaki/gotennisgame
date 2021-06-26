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
	ebiten.SetWindowTitle(g.gameName)
	g.botPaddle = newPaddle(float64(g.screenWidth)/2-float64(paddleWidth)/2, 0)
	g.playerPaddle = newPaddle(float64(g.screenWidth)/2-float64(paddleWidth)/2, float64(g.screenHeight)-float64(paddleHeight))
}

func (g *game) Update() error {
	cursorX, _ := ebiten.CursorPosition()
	windowWidth, _ := ebiten.WindowSize()
	newPlayerX := float64(cursorX) - float64(paddleWidth)/2
	if newPlayerX < 0 {
		newPlayerX = 0
	} else if newPlayerX+float64(paddleWidth) > float64(windowWidth) {
		newPlayerX = float64(windowWidth) - float64(paddleWidth)
	}
	g.playerPaddle.setX(newPlayerX)
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{})
	botPaddleOp := &ebiten.DrawImageOptions{}
	botPaddleOp.GeoM.Translate(g.botPaddle.getX(), g.botPaddle.getY())
	screen.DrawImage(g.botPaddle.Image, botPaddleOp)

	playerPaddleOp := &ebiten.DrawImageOptions{}
	playerPaddleOp.GeoM.Translate(g.playerPaddle.getX(), g.playerPaddle.getY())
	screen.DrawImage(g.playerPaddle.Image, playerPaddleOp)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
