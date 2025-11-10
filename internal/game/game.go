package game

import (
	"tictactoe_5x5/internal/graphics"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Board  *Board
	Assets *graphics.Assets
	Turn   rune
	Winner rune
}

func New() *Game {
	g := &Game{
		Board:  NewBoard(5, 5),
		Assets: graphics.LoadAssets(),
		Turn:   'W',
	}
	return g
}

func (g *Game) Update() error {
	if g.Winner != 0 {
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			g.Board.Clear()
			g.Winner = 0
			g.Turn = 'W'
		}
		return nil
	}

	x, y := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		bx, by := x/32, y/32
		if g.Board.Place(bx, by, g.Turn) {
			if CheckWin(g.Board, g.Turn) {
				g.Winner = g.Turn
			} else {
				if g.Turn == 'W' {
					g.Turn = 'B'
				} else {
					g.Turn = 'W'
				}
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Assets.DrawBoard(screen)
	g.Board.Draw(screen, g.Assets)

	if g.Winner != 0 {
		ebitenutil.DebugPrint(screen, string(g.Winner)+" wins! Press R to restart.")
	} else {
		ebitenutil.DebugPrint(screen, "Turn: "+string(g.Turn))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 32 * g.Board.W, 32 * g.Board.H
}
