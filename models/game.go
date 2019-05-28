package models

import (
	"fmt"
)

// Game - the play state for a game of tic-tac-toe, consisting of board state
//   and the mark currently in play.
type Game struct {
	CurrentMark rune       `json:"current_mark"`
	Board       [3][3]rune `json:"board"`
}

func (g Game) String() string {
	gamestr := " === Current Board ===\n\n"
	gamestr += fmt.Sprintf("       %c | %c | %c \n      -----------\n",
		g.Board[0][0], g.Board[0][1], g.Board[0][2])
	gamestr += fmt.Sprintf("       %c | %c | %c \n      -----------\n",
		g.Board[1][0], g.Board[1][1], g.Board[1][2])
	gamestr += fmt.Sprintf("       %c | %c | %c \n\n",
		g.Board[2][0], g.Board[2][1], g.Board[2][2])
	gamestr += fmt.Sprintf("Current player: '%c'\n", g.CurrentMark)
	return gamestr
}
