package main

func checkVictoryLines(game Game) bool {
	symbol := game.players[game.current]
	for row := 0; row < rows; row++ {
		count := 0
		for col := 0; col < columns; col++ {
			if game.board[row][col] == symbol {
				count++
				if count == 4 {
					return true
				}
			} else {
				count = 0
			}
		}
	}
	return false
}

func checkVictoryColumns(game Game) bool {
	symbol := game.players[game.current]
	for col := 0; col < columns; col++ {
		count := 0
		for row := 0; row < rows; row++ {
			if game.board[row][col] == symbol {
				count++
				if count == 4 {
					return true
				}
			} else {
				count = 0
			}
		}
	}
	return false
}

func checkVictoryDiagonals(game Game) bool {
	symbol := game.players[game.current]
	// Diagonale descendante ()
	for row := 0; row < rows-3; row++ {
		for col := 0; col < columns-3; col++ {
			if game.board[row][col] == symbol &&
				game.board[row+1][col+1] == symbol &&
				game.board[row+2][col+2] == symbol &&
				game.board[row+3][col+3] == symbol {
				return true
			}
		}
	}
	// Diagonale montante (/)
	for row := 3; row < rows; row++ {
		for col := 0; col < columns-3; col++ {
			if game.board[row][col] == symbol &&
				game.board[row-1][col+1] == symbol &&
				game.board[row-2][col+2] == symbol &&
				game.board[row-3][col+3] == symbol {
				return true
			}
		}
	}
	return false
}

func checkDraw(game Game) bool {
	for col := 0; col < columns; col++ {
		if game.board[0][col] == empty {
			return false
		}
	}
	return true
}
