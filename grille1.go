package main

func displayBoard(game *Game) {
    fmt.Println("\nÉtat actuel de la grille :")
    for col := 0; col < columns; col++ {
        fmt.Printf(" %d", col)
    }
    fmt.Println()
    for , row := range game.board {
        for , cell := range row {
            fmt.Printf(" %c", cell)
        }
        fmt.Println()
    }
    fmt.Println()
}