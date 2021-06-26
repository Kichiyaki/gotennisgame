package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"

	"github.com/Kichiyaki/gotennisgame/game"
)

const (
	screenWidth  = 800
	screenHeight = 600
	gameName     = "Tennis Game"
)

func main() {
	g, err := game.New(game.Config{
		WindowWidth:  screenWidth,
		WindowHeight: screenHeight,
		GameName:     gameName,
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
