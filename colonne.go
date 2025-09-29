package main

import "fmt"

func dropDisc(game *Game, col int) bool {
	if col < 0 || col >= columns {
		fmt.Println("Colonne invalide. Essayez de nouveau.")
		return false
	}
	for row := rows - 1; row >= 0; row-- {
		if game.board[row][col] == empty {
			game.board[row][col] = game.players[game.current]
			return true
		}
	}
	fmt.Println("Colonne pleine. Choisissez une autre colonne.")
	return false
}
