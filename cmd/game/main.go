package main

import (
	"log"

	"tictactoe_5x5/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Tic Tac Toe")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
