package smallest_test

import (
	"github.com/krzysztofSkolimowski/programming-practice/cracking-6ed/16-moderate/smallest-difference"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallestDifference(t *testing.T) {
	testCases := []struct {
		name         string
		a, b         []int
		expectedDiff int
	}{
		{
			name:         "",
			a:            []int{0},
			b:            []int{0},
			expectedDiff: 0,
		},
		{
			name:         "",
			a:            []int{0, 1, 2, 3, 4},
			b:            []int{10, 9, 8, 7, 6, 5},
			expectedDiff: 1,
		},
		{
			name:         "",
			a:            []int{3, 4, 5, 6},
			b:            []int{10, 9, 8, 7, 6, 5},
			expectedDiff: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualDiff := smallest.Difference(tc.a, tc.b)
			assert.Equal(t, tc.expectedDiff, actualDiff)
		})
	}
}

func TestAbs(t *testing.T) {
	testCases := [] struct {
		name        string
		x, expected int
	}{
		{
			name:     "",
			x:        1,
			expected: 1,
		},
		{
			name:     "",
			x:        -1,
			expected: 1,
		},
		{
			name:     "",
			x:        0,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, smallest.Abs(tc.x))
		})
	}
}
