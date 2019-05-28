package controllers

import (
	"errors"

	"github.com/kbrimm/tic-tac-toe/models"
)

// StartNewGame - initalizes new board state for a game of tic-tac-toe
func StartNewGame(mark rune) models.Game {
	var newBoard [3][3]rune
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newBoard[i][j] = ' '
		}
	}
	return models.Game{CurrentMark: mark, Board: newBoard}
}

// PlaceMark - given a game configuration and coordinates, adds the mark,
//   updates the current player and returns the new board state. Returns error
//   if the space is invalid or already occupied
func PlaceMark(game models.Game, row int, col int) (models.Game, error) {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return game, errors.New("gameplay error: invalid index")
	}

	if game.Board[row][col] != ' ' {
		return game, errors.New("gameplay error: space already occupied")
	}

	game.Board[row][col] = game.CurrentMark
	game.CurrentMark = switchCurrentMark(game.CurrentMark)
	return game, nil
}

func switchCurrentMark(mark rune) rune {
	if mark == 'x' {
		return 'o'
	}
	return 'x'
}
