package game

import (
	"tictactoe_5x5/internal/graphics"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	W, H  int
	Cells [][]rune
}

func NewBoard(w, h int) *Board {
	cells := make([][]rune, h)
	for i := range cells {
		cells[i] = make([]rune, w)
	}
	return &Board{W: w, H: h, Cells: cells}
}

func (b *Board) Clear() {
	for y := range b.Cells {
		for x := range b.Cells[y] {
			b.Cells[y][x] = 0
		}
	}
}

func (b *Board) Place(x, y int, mark rune) bool {
	if x < 0 || y < 0 || x >= b.W || y >= b.H {
		return false
	}
	if b.Cells[y][x] != 0 {
		return false
	}
	b.Cells[y][x] = mark
	return true
}

func (b *Board) Draw(screen *ebiten.Image, a *graphics.Assets) {
	for y := 0; y < b.H; y++ {
		for x := 0; x < b.W; x++ {
			switch b.Cells[y][x] {
			case 'W':
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*32+4), float64(y*32+4))
				screen.DrawImage(a.WhitePlayer, op)
			case 'B':
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*32+4), float64(y*32+4))
				screen.DrawImage(a.BlackPlayer, op)
			}
		}
	}
}
