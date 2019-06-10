package trailling_zeroes_test

import (
	"github.com/krzysztofSkolimowski/programming-practice/cracking-the-coding-interview-6th-ed/chapter-16-moderate/trailling_zeroes"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCountTraillingZeroes(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			n:        2,
			expected: 0,
		},
		{
			n:        5,
			expected: 1,
		},
		{
			n:        20,
			expected: 4,
		},
		{
			n:        40,
			expected: 9,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, trailling_zeroes.CountTrailingZeroesFact(tc.n), tc.expected)
		})
	}
}
