package number_swapper_test

import (
	"github.com/krzysztofSkolimowski/programming-practice/cracking-the-coding-interview-6th-ed/chapter-16-moderate/number_swapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	tests := []struct {
		name string
		a, b int
	}{
		{
			name: "different_positive_numbers_correctly_swapped_lower_first",
			a:    1,
			b:    2,
		},
		{
			name: "different_positive_numbers_correctly_swapped_higher_first",
			a:    2,
			b:    1,
		},
		{
			name: "same_number_does_not_change",
			a:    1,
			b:    1,
		},
		{
			name: "negative_numbers_swapped_correctly",
			a:    -1,
			b:    -2,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			expectedA, expectedB := tc.b, tc.a
			number_swapper.Swap(&tc.a, &tc.b)
			assert.Equal(t, expectedA, tc.a)
			assert.Equal(t, expectedB, tc.b)
		})
	}
}
