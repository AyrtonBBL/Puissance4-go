package main

const (
	rows    = 6
	columns = 7
	empty   = '.'
)

type Game struct {
	board    [rows][columns]rune
	players  [2]rune
	current  int
	turns    int
	gameOver bool
}
