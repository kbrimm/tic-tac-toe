package controllers

import (
	"testing"

	"github.com/kbrimm/tic-tac-toe/models"
	"github.com/stretchr/testify/assert"
)

func TestStartNewGame(t *testing.T) {
	newGame := StartNewGame('x')

	expectedBoard := [3][3]rune{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '}}
	assert.Equal(t, 'x', newGame.CurrentMark, "current mark is not 'x'")
	assert.Equal(t, expectedBoard, newGame.Board, "new board is not 3x3 of ' '")
}

func TestPlaceMark(t *testing.T) {
	game, err := PlaceMark(StartNewGame('x'), 0, 0)
	assert.Equal(t, nil, err, "error when placing mark")
	assert.Equal(t, 'x', game.Board[0][0], "correct mark not found")
	assert.Equal(t, 'o', game.CurrentMark, "current mark not updated")

	game, err = PlaceMark(game, 0, 0)
	assert.NotEqual(t, nil, err, "no error raised when placing double mark")
	assert.Equal(t, 'x', game.Board[0][0], "correct mark not left on board")
	assert.Equal(t, 'o', game.CurrentMark, "current mark updated despite invalid move")

	game, err = PlaceMark(game, 0, 3)
	assert.NotEqual(t, nil, err, "no error raised when mark out of bounds")
}

func TestScoreGame(t *testing.T) {
	board := [3][3]rune{
		{'o', 'o', 'x'},
		{'o', 'x', 'x'},
		{' ', 'o', ' '}}
	game := models.Game{Board: board, CurrentMark: 'x'}
	hasStalemate, hasWin, winningMark := ScoreGame(game)
	assert.False(t, hasStalemate, "stalemate condition detected incorrectly")
	assert.False(t, hasWin, "win condition detected incorrectly")
	assert.Equal(t, ' ', winningMark, "winner declared with no win condition")

	board = [3][3]rune{
		{'o', 'o', 'x'},
		{'x', 'x', 'x'},
		{' ', 'o', ' '}}
	game = models.Game{Board: board, CurrentMark: 'x'}
	_, hasWin, winningMark = ScoreGame(game)
	assert.True(t, hasWin, "row win condition not detected")
	assert.Equal(t, 'x', winningMark, "incorrect row winner declared")

	board = [3][3]rune{
		{'o', 'o', 'x'},
		{'x', 'o', 'x'},
		{' ', 'o', ' '}}
	game = models.Game{Board: board, CurrentMark: 'x'}
	_, hasWin, winningMark = ScoreGame(game)
	assert.True(t, hasWin, "column win condition not detected")
	assert.Equal(t, 'o', winningMark, "incorrect column winner declared")

	board = [3][3]rune{
		{'o', ' ', 'x'},
		{' ', 'o', 'x'},
		{' ', ' ', 'o'}}
	game = models.Game{Board: board, CurrentMark: 'x'}
	_, hasWin, winningMark = ScoreGame(game)
	assert.True(t, hasWin, "diagonal win condition not detected")
	assert.Equal(t, 'o', winningMark, "incorrect diagonal winner declared")

	board = [3][3]rune{
		{'o', 'x', 'o'},
		{'o', 'x', ' '},
		{'x', 'o', 'x'}}
	game = models.Game{Board: board, CurrentMark: 'x'}
	hasStalemate, _, _ = ScoreGame(game)
	assert.True(t, hasWin, "stalemate condition not detected")
}
