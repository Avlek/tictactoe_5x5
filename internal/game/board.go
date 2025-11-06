package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Board struct {
	boardImage *ebiten.Image
}

// NewBoard generates a new Board with giving a size.
func NewBoard() (*Board, error) {
	boardImage, _, err := ebitenutil.NewImageFromFile("internal/assets/images/board.png")
	if err != nil {
		return nil, err
	}
	return &Board{boardImage: boardImage}, nil
}

func (b *Board) Update() {}

func (b *Board) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(20, 20)

	screen.DrawImage(b.boardImage, &opts)
}
