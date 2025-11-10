package game

func CheckWin(b *Board, mark rune) bool {
	for y := 0; y < b.H; y++ {
		for x := 0; x < b.W; x++ {
			if b.Cells[y][x] != mark {
				continue
			}

			if x <= b.W-4 &&
				b.Cells[y][x+1] == mark &&
				b.Cells[y][x+2] == mark &&
				b.Cells[y][x+3] == mark {
				return true
			}

			if y <= b.H-4 &&
				b.Cells[y+1][x] == mark &&
				b.Cells[y+2][x] == mark &&
				b.Cells[y+3][x] == mark {
				return true
			}

			if x <= b.W-4 && y <= b.H-4 &&
				b.Cells[y+1][x+1] == mark &&
				b.Cells[y+2][x+2] == mark &&
				b.Cells[y+3][x+3] == mark {
				return true
			}

			if x >= 3 && y <= b.H-4 &&
				b.Cells[y+1][x-1] == mark &&
				b.Cells[y+2][x-2] == mark &&
				b.Cells[y+3][x-3] == mark {
				return true
			}
		}
	}

	return false
}
