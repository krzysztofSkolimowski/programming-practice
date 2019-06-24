package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Won(t *testing.T) {
	tests := []struct {
		name                 string
		lastMoveX, lastMoveY int
		v                    Value
		g                    game
		expected             bool
	}{
		{
			name:      "not_win",
			lastMoveX: 0,
			lastMoveY: 1,
			v:         X,
			g: game{
				[3][3]Value{
					{X, X, Empty},
					{Empty, Empty, Empty},
					{O, O, Empty},
				},
			},
			expected: false,
		},
		{
			name:      "horizontal_X_won",
			lastMoveX: 0,
			lastMoveY: 0,
			v:         X,
			g: game{
				[3][3]Value{
					{X, X, X},
					{Empty, Empty, Empty},
					{O, O, Empty},
				},
			},
			expected: true,
		},
		{
			name:      "vertical_X_won",
			lastMoveX: 0,
			lastMoveY: 0,
			v:         X,
			g: game{
				[3][3]Value{
					{X, O, X},
					{X, Empty, O},
					{X, O, Empty},
				},
			},
			expected: true,
		},
		{
			name:      "diagonal_top_X_won",
			lastMoveX: 1,
			lastMoveY: 1,
			v:         X,
			g: game{
				[3][3]Value{
					{X, O, X},
					{O, X, Empty},
					{O, O, X},
				},
			},
			expected: true,
		},
		{
			name:      "diagonal_bottom_X_won",
			lastMoveX: 1,
			lastMoveY: 1,
			v:         X,
			g: game{
				[3][3]Value{
					{O, O, X},
					{Empty, X, Empty},
					{X, Empty, Empty},
				},
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.g.Won(tc.v, tc.lastMoveX, tc.lastMoveY))
		})
	}

}
