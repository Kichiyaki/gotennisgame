package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pkg/errors"
	"image/color"
)

type Config struct {
	GameName     string
	WindowWidth  int
	WindowHeight int
}

type game struct {
	gameName     string
	windowWidth  int
	windowHeight int
	playerPaddle *paddle
	botPaddle    *paddle
	ball         *ball
}

func New(cfg Config) (ebiten.Game, error) {
	g := &game{
		windowHeight: cfg.WindowHeight,
		windowWidth:  cfg.WindowWidth,
		gameName:     cfg.GameName,
	}
	err := g.init()
	return g, err
}

func (g *game) init() error {
	var err error

	ebiten.SetWindowSize(g.windowWidth, g.windowHeight)
	ebiten.SetWindowTitle(g.gameName)

	g.botPaddle = newPaddle(float64(g.windowWidth)/2-float64(paddleWidth)/2, 0)
	g.playerPaddle = newPaddle(float64(g.windowWidth)/2-float64(paddleWidth)/2, float64(g.windowHeight)-float64(paddleHeight))

	g.ball, err = newBall(float64(g.windowWidth)/2, float64(g.windowHeight)/2)
	if err != nil {
		return errors.Wrap(err, "couldn't initialize the game")
	}

	return nil
}

func (g *game) Update() error {
	g.updatePlayerPosition()
	g.updateBallPosition()
	return nil
}

func (g *game) updatePlayerPosition() {
	windowWidth, _ := ebiten.WindowSize()
	cursorX, _ := ebiten.CursorPosition()
	newPlayerX := float64(cursorX) - float64(paddleWidth)/2
	if newPlayerX < 0 {
		newPlayerX = 0
	} else if newPlayerX+float64(paddleWidth) > float64(windowWidth) {
		newPlayerX = float64(windowWidth) - float64(paddleWidth)
	}
	g.playerPaddle.setX(newPlayerX)
}

func (g *game) updateBallPosition() {
	windowWidth, windowHeight := ebiten.WindowSize()
	ballWidth, ballHeight := g.ball.Size()
	ballX := g.ball.getX()
	centerBallX := ballX + float64(ballWidth)/2
	ballY := g.ball.getY()
	if g.playerPaddle.isPointInBoundaries(centerBallX, ballY+float64(ballHeight)) || g.botPaddle.isPointInBoundaries(centerBallX, ballY) {
		g.ball.velocity.setX(g.ball.velocity.getX() * -1)
		g.ball.velocity.setY(g.ball.velocity.getY() * -1)
	} else if ballX < 0 || ballX+float64(ballWidth) > float64(windowWidth) {
		g.ball.velocity.setX(g.ball.velocity.getX() * -1)
	} else if ballY < 0 || ballY+float64(ballHeight) > float64(windowHeight) {
		g.ball.velocity.setY(g.ball.velocity.getY() * -1)
	}
	g.ball.setX(ballX + g.ball.velocity.getX())
	g.ball.setY(ballY + g.ball.velocity.getY())
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
	return g.windowWidth, g.windowHeight
}
