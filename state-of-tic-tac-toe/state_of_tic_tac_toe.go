package stateoftictactoe

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

var directions = []struct {
	x int
	y int
	// }{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
}{{0, 1}, {1, -1}, {1, 0}, {1, 1}}

func StateOfTicTacToe(board []string) (State, error) {
	for i, row := range board {
		for j := range row {
			for _, d := range directions {
				x, y := i, j
				x += d.x
				y += d.y

				if x >= 3 || y >= 3 || x < 0 || y < 0 || x+d.x >= 3 || x+d.x < 0 || y+d.y >= 3 || y-d.y < 0 {
					continue
				}

				if (board[i][j] == 'X' && board[x][y] == 'X' && board[x+d.x][y+d.y] == 'X') ||
					board[i][j] == 'O' && board[x][y] == 'O' && board[x+d.x][y+d.y] == 'O' {
					return Win, nil
				}
			}
		}
	}
	return Ongoing, nil
}
