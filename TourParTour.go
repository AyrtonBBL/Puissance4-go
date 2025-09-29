package main

func switchPlayer(game *Game) {
	game.current = 1 - game.current
	game.turns++
}
