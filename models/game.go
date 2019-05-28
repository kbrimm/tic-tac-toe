package models

// Game - the play state for a game of tic-tac-toe, consisting of board state
//   and the mark currently in play.
type Game struct {
	CurrentMark rune
	Board       [3][3]rune
}
