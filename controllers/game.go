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

// PlaceMark - given a game configuration and coordinates, adds mark,
//   updates the current player
func PlaceMark(game models.Game, row int, col int) (newGameState models.Game, err error) {
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

// ScoreGame - evaluates the current board for end-game conditions
func ScoreGame(game models.Game) (stalemate bool, win bool, winner rune) {
	if checkWins('x', game.Board) {
		return false, true, 'x'
	}

	if checkWins('o', game.Board) {
		return false, true, 'o'
	}

	return checkStalemate(game.Board), false, ' '
}

func checkWins(mark rune, board [3][3]rune) bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if mark == board[i][0] && mark == board[i][1] && mark == board[i][2] {
			return true
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if mark == board[0][i] && mark == board[1][i] && mark == board[2][i] {
			return true
		}
	}

	// Check diagonals
	if (mark == board[0][0] && mark == board[1][1] && mark == board[2][2]) ||
		(mark == board[0][2] && mark == board[1][1] && mark == board[2][0]) {
		return true
	}
	return false
}

func checkStalemate(board [3][3]rune) bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if notBlocked(board[i][0], board[i][1]) &&
			notBlocked(board[i][1], board[i][2]) &&
			notBlocked(board[i][0], board[i][2]) {
			return false
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if notBlocked(board[0][i], board[1][i]) &&
			notBlocked(board[1][i], board[2][i]) &&
			notBlocked(board[0][i], board[2][i]) {
			return false
		}
	}

	// Check diagonals
	if notBlocked(board[0][0], board[1][1]) &&
		notBlocked(board[1][1], board[2][2]) &&
		notBlocked(board[0][0], board[2][2]) {
		return false
	}

	if notBlocked(board[0][2], board[1][1]) &&
		notBlocked(board[0][2], board[2][0]) &&
		notBlocked(board[1][1], board[0][2]) {
		return false
	}
	return true
}

func notBlocked(markOne rune, markTwo rune) bool {
	return markOne == ' ' || markTwo == ' ' || markOne == markTwo
}
