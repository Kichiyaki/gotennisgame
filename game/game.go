package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pkg/errors"
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
	ball         *ball
}

func New(cfg Config) (ebiten.Game, error) {
	g := &game{
		screenHeight: cfg.ScreenHeight,
		screenWidth:  cfg.ScreenWidth,
		gameName:     cfg.GameName,
	}
	err := g.init()
	return g, err
}

func (g *game) init() error {
	var err error

	ebiten.SetWindowSize(g.screenWidth, g.screenHeight)
	ebiten.SetWindowTitle(g.gameName)

	g.botPaddle = newPaddle(float64(g.screenWidth)/2-float64(paddleWidth)/2, 0)
	g.playerPaddle = newPaddle(float64(g.screenWidth)/2-float64(paddleWidth)/2, float64(g.screenHeight)-float64(paddleHeight))

	g.ball, err = newBall(float64(g.screenWidth)/2, float64(g.screenHeight)/2)
	if err != nil {
		return errors.Wrap(err, "couldn't initialize the game")
	}

	return nil
}

func (g *game) Update() error {
	windowWidth, _ := ebiten.WindowSize()

	cursorX, _ := ebiten.CursorPosition()
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

	ballOp := &ebiten.DrawImageOptions{}
	ballOp.GeoM.Translate(g.ball.getX(), g.ball.getY())
	screen.DrawImage(g.ball.Image, ballOp)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
