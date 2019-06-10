package tictactoe_test

import (
	"testing"

	"github.com/krzysztofSkolimowski/programming-practice/cracking-the-coding-interview-6th-ed/chapter-16-moderate/tictactoe"
	"github.com/stretchr/testify/assert"
)

func Test_game_Move(t *testing.T) {
	tests := []struct {
		name         string
		x, y         int
		v, expectedV tictactoe.Value
		expectedErr  error
	}{
		{
			x:           0,
			y:           0,
			v:           tictactoe.X,
			expectedV:   tictactoe.X,
			expectedErr: nil,
		},
		{
			x:           0,
			y:           0,
			v:           tictactoe.O,
			expectedV:   tictactoe.O,
			expectedErr: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := tictactoe.NewGame()
			_, err := g.Move(tc.x, tc.y, tc.v)

			assert.Equal(t, tc.expectedErr, err)

			assert.Equal(t, tc.expectedV, g.Field(tc.x, tc.y))
		})
	}
}

func Test_game_Move_Err_Field_Is_Empty(t *testing.T) {
	tests := []struct {
		name          string
		x, y          int
		firstV, nextV tictactoe.Value
		expectedErr   error
	}{
		{
			x:           0,
			y:           0,
			firstV:      tictactoe.O,
			nextV:       tictactoe.O,
			expectedErr: tictactoe.ErrFieldNotEmpty,
		},
		{
			x:           0,
			y:           0,
			firstV:      tictactoe.X,
			nextV:       tictactoe.X,
			expectedErr: tictactoe.ErrFieldNotEmpty,
		},
		{
			x:           0,
			y:           0,
			firstV:      tictactoe.X,
			nextV:       tictactoe.O,
			expectedErr: tictactoe.ErrFieldNotEmpty,
		},
		{
			x:           0,
			y:           0,
			firstV:      tictactoe.O,
			nextV:       tictactoe.X,
			expectedErr: tictactoe.ErrFieldNotEmpty,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := tictactoe.NewGame()
			_, err := g.Move(tc.x, tc.y, tc.firstV)
			_, err = g.Move(tc.x, tc.y, tc.nextV)

			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.firstV, g.Field(tc.x, tc.y))
		})
	}
}
