package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/pkg/errors"
	"image/color"

	"github.com/Kichiyaki/gotennisgame/game/asset"
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
	score        *score
	font         asset.Font
}

func New(cfg Config) (ebiten.Game, error) {
	g := &game{
		windowHeight: cfg.WindowHeight,
		windowWidth:  cfg.WindowWidth,
		gameName:     cfg.GameName,
		score:        &score{},
	}
	err := g.init()
	return g, err
}

func (g *game) init() error {
	var err error

	g.font, err = asset.NewFont()
	if err != nil {
		return errors.Wrap(err, "couldn't initialize the game")
	}

	ebiten.SetWindowSize(g.windowWidth, g.windowHeight)
	ebiten.SetWindowTitle(g.gameName)

	paddleWidth := g.windowWidth / 6
	paddleHeight := g.windowHeight / 32
	g.botPaddle = newPaddle(float64(g.windowWidth)/2-float64(paddleWidth)/2, 0, paddleWidth, paddleHeight)
	g.playerPaddle = newPaddle(float64(g.windowWidth)/2-float64(paddleWidth)/2, float64(g.windowHeight)-float64(paddleHeight), paddleWidth, paddleHeight)

	g.ball, err = newBall(float64(g.windowWidth)/2, float64(g.windowHeight)/2)
	if err != nil {
		return errors.Wrap(err, "couldn't initialize the game")
	}

	return nil
}

func (g *game) Update() error {
	g.updatePlayerPaddlePosition()
	g.updateBotPaddlePosition()
	g.updateBallPosition()
	return nil
}

func (g *game) updatePlayerPaddlePosition() {
	windowWidth, _ := ebiten.WindowSize()
	cursorX, _ := ebiten.CursorPosition()
	paddleWidth, _ := g.playerPaddle.Size()
	newPlayerX := float64(cursorX) - float64(paddleWidth)/2
	if newPlayerX < 0 {
		newPlayerX = 0
	} else if newPlayerX+float64(paddleWidth) > float64(windowWidth) {
		newPlayerX = float64(windowWidth) - float64(paddleWidth)
	}
	g.playerPaddle.setX(newPlayerX)
}

func (g *game) updateBotPaddlePosition() {
	windowWidth, _ := ebiten.WindowSize()
	newBotX := g.botPaddle.getX() + g.botPaddle.speed.getX()
	paddleWidth, _ := g.playerPaddle.Size()
	if newBotX > g.ball.getX() {
		newBotX -= g.botPaddle.speed.getX() * 2
	}
	if newBotX < 0 {
		newBotX = 0
	} else if newBotX+float64(paddleWidth) > float64(windowWidth) {
		newBotX = float64(windowWidth) - float64(paddleWidth)
	}
	g.botPaddle.setX(newBotX)
}

func (g *game) updateBallPosition() {
	windowWidth, windowHeight := ebiten.WindowSize()
	resetPosition := false
	_, paddleHeight := g.playerPaddle.Size()
	if g.playerPaddle.isPointInBoundaries(g.ball.getMidX(), g.ball.getBottomY()) || g.botPaddle.isPointInBoundaries(g.ball.getMidX(), g.ball.getY()) {
		g.ball.speed.setY(g.ball.speed.getY() * -1)
	} else if g.ball.getX() < 0 || g.ball.getRightX() > float64(windowWidth) {
		g.ball.speed.setX(g.ball.speed.getX() * -1)
	} else if g.ball.getY() < float64(paddleHeight) {
		g.score.addToPlayerScore(1)
		resetPosition = true
	} else if g.ball.getBottomY() > float64(windowHeight)-float64(paddleHeight) {
		g.score.addToBotScore(1)
		resetPosition = true
	}

	if resetPosition {
		g.ball.resetPosition()
		g.botPaddle.resetPosition()
		g.playerPaddle.resetPosition()
	}

	g.ball.setX(g.ball.getX() + g.ball.speed.getX())
	g.ball.setY(g.ball.getY() + g.ball.speed.getY())
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	g.drawNet(screen)
	g.drawScore(screen)

	for _, toRender := range []renderable{
		g.botPaddle,
		g.playerPaddle,
		g.ball,
	} {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(toRender.getX(), toRender.getY())
		screen.DrawImage(toRender.getImage(), op)
	}
}

func (g *game) drawScore(screen *ebiten.Image) {
	_, windowHeight := ebiten.WindowSize()
	normalFont := g.font.GetNormal()
	uppercaseHeight := normalFont.Metrics().CapHeight.Floor() * -1
	text.Draw(screen, fmt.Sprintf("Bot: %d", g.score.getBotScore()), normalFont, 10, uppercaseHeight*2, color.White)
	text.Draw(screen, fmt.Sprintf("Player: %d", g.score.getPlayerScore()), normalFont, 10, windowHeight-uppercaseHeight, color.White)
}

func (g *game) drawNet(screen *ebiten.Image) {
	windowWidth, windowHeight := ebiten.WindowSize()
	rectWidth := float64(windowWidth) / 19
	rectHeight := float64(windowHeight) / 50

	for i := 0.0; i <= float64(windowWidth); i += rectWidth * 2 {
		ebitenutil.DrawRect(screen, i, float64(windowHeight)/2-float64(rectHeight)/2, rectWidth, rectHeight, color.White)
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.windowWidth, g.windowHeight
}
