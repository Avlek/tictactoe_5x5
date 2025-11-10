package graphics

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Assets struct {
	Board       *ebiten.Image
	WhitePlayer *ebiten.Image
	BlackPlayer *ebiten.Image
}

func LoadAssets() *Assets {
	board := loadImage("internal/assets/images/board.png")
	white := loadImage("internal/assets/images/white_player.png")
	black := loadImage("internal/assets/images/black_player.png")

	return &Assets{Board: board, WhitePlayer: white, BlackPlayer: black}
}

func loadImage(path string) *ebiten.Image {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func (a *Assets) DrawBoard(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(a.Board, op)
}
