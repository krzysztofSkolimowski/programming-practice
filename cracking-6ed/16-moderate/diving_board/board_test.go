package diving_board

import (
	"reflect"
	"testing"
)

func Test_getAllLenghts(t *testing.T) {
	const shorterBoardLenght int = 1
	const longerBoardLenght int = 100

	tests := []struct {
		name           string
		numberOfBoards int
		expected       map[int]struct{}
	}{
		{
			name:           "one_board",
			numberOfBoards: 1,
			expected: map[int]struct{}{
				1:   {},
				100: {},
			},
		},
		{
			name:           "two_boards",
			numberOfBoards: 2,
			expected: map[int]struct{}{
				1:   {},
				2:   {},
				100: {},
				101: {},
				200: {},
			},
		},
		{
			name:           "three_boards",
			numberOfBoards: 3,
			expected: map[int]struct{}{
				1:   {},
				2:   {},
				3:   {},
				100: {},
				101: {},
				102: {},
				200: {},
				201: {},
				300: {},
			},
		},
		{
			name:           "many_boards",
			numberOfBoards: 5,
			expected: map[int]struct{}{
				1:   {},
				2:   {},
				3:   {},
				4:   {},
				5:   {},
				100: {},
				101: {},
				102: {},
				103: {},
				104: {},
				200: {},
				201: {},
				202: {},
				203: {},
				300: {},
				301: {},
				302: {},
				400: {},
				401: {},
				500: {},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getAllLenghts(tc.numberOfBoards, shorterBoardLenght, longerBoardLenght)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("getAllLenghts() = %v, actual %v", actual, tc.expected)
			}
		})
	}
}

func Test_getAllLenghtsSums(t *testing.T) {
	const shorterBoardLenght int = 1
	const longerBoardLenght int = 100

	tests := []struct {
		name           string
		numberOfBoards int
		expected       map[int]struct{}
	}{
		{
			name:           "one_board",
			numberOfBoards: 1,
			expected: map[int]struct{}{
				1:   {},
				100: {},
			},
		},
		{
			name:           "two_boards",
			numberOfBoards: 2,
			expected: map[int]struct{}{
				2:   {},
				101: {},
				200: {},
			},
		},
		{
			name:           "three_boards",
			numberOfBoards: 3,
			expected: map[int]struct{}{
				3:   {},
				102: {},
				201: {},
				300: {},
			},
		},
		{
			name:           "many_boards",
			numberOfBoards: 5,
			expected: map[int]struct{}{
				5:   {},
				104: {},
				203: {},
				302: {},
				401: {},
				500: {},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getAllLenghtsSums(tc.numberOfBoards, shorterBoardLenght, longerBoardLenght)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("getAllLenghts() = %v, actual %v", actual, tc.expected)
			}
		})
	}
}
