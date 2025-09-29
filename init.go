package main

func main() {
    game := initializeGame()

    for !game.gameOver {
        displayBoard(&game)
        fmt.Printf("Tour du joueur %c. Choisissez une colonne (0 à %d) : ", game.players[game.current], columns-1)

        var col int
        _, err := fmt.Scan(&col)
        if err != nil {
            fmt.Println("Entrée invalide. Veuillez entrer un numéro de colonne.")
            continue
        }

        if !dropDisc(&game, col) {
            continue
        }

        if checkVictoryLines(&game)  checkVictoryColumns(&game)  checkVictoryDiagonals(&game) {
            displayBoard(&game)
            fmt.Printf("Le joueur %c a gagné !\n", game.players[game.current])
            game.gameOver = true
            break
        }

        if checkDraw(&game) {
            displayBoard(&game)
            fmt.Println("Match nul !")
            game.gameOver = true
            break
        }

        switchPlayer(&game)
    }
}

func initializeGame() Game {
    var game Game
    for i := range game.board {
        for j := range game.board[i] {
            game.board[i][j] = empty
        }
    }
    game.players[0] = 'X'
    game.players[1] = 'O'
    game.current = 0
    game.turns = 0
    game.gameOver = false
    return game
} 
