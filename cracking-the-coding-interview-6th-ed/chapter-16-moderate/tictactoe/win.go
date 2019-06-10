package tictactoe

import (
	"github.com/pkg/errors"
)

var ErrFieldNotEmpty = errors.New("field is not empty")

// dimension
const dim = 3

type Value int

const (
	Empty = iota
	X
	O
)

func (v Value) String() string { return [...]string{"Empty", "X", "O"}[v] }

type game struct {
	board [dim][dim]Value
}

func NewGame() *game {
	board := [dim][dim]Value{}

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			board[i][j] = Empty
		}
	}

	return &game{board}
}

func (g game) Field(x, y int) Value { return g.board[x][y] }

func (g *game) Move(x, y int, v Value) (bool, error) {
	if g.board[x][y] != Empty {
		return false, ErrFieldNotEmpty
	}
	g.board[x][y] = v
	return g.Won(v, x, y), nil
}

func (g game) Won(value Value, x, y int) bool {
	// check horizontal
	sum := 0
	for i := 0; i < dim; i++ {
		if g.board[i][y] == value {
			sum++
		} else {
			break
		}

		if sum == dim {
			return true
		}
	}

	// check vertical
	sum = 0
	for i := 0; i < dim; i++ {
		if g.board[x][i] == value {
			sum++
		} else {
			break
		}
		if sum == dim {
			return true
		}
	}

	// check diagonal
	onDiagonal := x == y || x+y == dim-1
	if onDiagonal {
		sum = 0
		// top-left to bottom-right
		for i := 0; i < dim; i++ {
			if g.board[i][i] == value {
				sum++
			} else {
				break
			}
			if sum == dim {
				return true
			}
		}

		// bottom-left to top-right
		for i := 0; i < dim; i++ {
			if g.board[dim-1-i][i] == value {
				sum++
			} else {
				break
			}
			if sum == dim {
				return true
			}
		}
	}

	return false
}
