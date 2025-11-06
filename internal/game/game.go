package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	input *Input
	board *Board
	turn  rune
}

func NewGame() (*Game, error) {
	board, err := NewBoard()
	if err != nil {
		return nil, err
	}
	return &Game{
		input: NewInput(),
		board: board,
		turn:  'X',
	}, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 200, 200
}
