package controllers

import (
	"testing"

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
