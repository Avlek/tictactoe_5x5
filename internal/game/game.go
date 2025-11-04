package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	board [3][3]rune
	turn  rune
}

func NewGame() *Game {
	return &Game{turn: 'X'}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Tic Tac Toe\n(X/O)")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 600, 600
}
