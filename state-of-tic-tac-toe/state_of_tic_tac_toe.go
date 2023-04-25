package stateoftictactoe

import "errors"

type State string

const Win, Draw, Ongoing State = "win", "draw", "ongoing"

func getPlayerStats(pl byte, board []string) (count int, won bool) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == pl {
				count++
			}
		}
	}
	for i := 0; i < 3; i++ {
		if board[i][0] == pl && board[i][1] == pl && board[i][2] == pl ||
			board[0][i] == pl && board[1][i] == pl && board[2][i] == pl {
			return count, true
		}
	}
	return count, board[0][0] == pl && board[1][1] == pl && board[2][2] == pl || board[0][2] == pl && board[1][1] == pl && board[2][0] == pl
}

func StateOfTicTacToe(board []string) (State, error) {
	Xcount, XWon := getPlayerStats('X', board)
	Ocount, OWon := getPlayerStats('O', board)

	if Xcount != Ocount && Xcount != Ocount+1 || XWon && OWon {
		return "", errors.New("invalid")
	} else if XWon || OWon {
		return Win, nil
	} else if Xcount+Ocount == 9 {
		return Draw, nil
	}
	return Ongoing, nil

}
