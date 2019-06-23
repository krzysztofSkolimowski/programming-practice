package english_int_test

import (
	"github.com/krzysztofSkolimowski/programming-practice/cracking-the-coding-interview-6th-ed/chapter-16-moderate/english_int"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name          string
		n             int
		expected      string
		expectedError error
	}{
		{
			name:     "zero",
			n:        0,
			expected: "zero",
		},
		{
			name:     "ones",
			n:        5,
			expected: "five",
		},
		{
			name:     "ten",
			n:        10,
			expected: "ten",
		},
		{
			name:     "eleven_to_nineteen",
			n:        16,
			expected: "sixteen",
		},
		{
			name:     "tens",
			n:        20,
			expected: "twenty",
		},
		{
			name:     "tens",
			n:        30,
			expected: "thirty",
		},
		{
			name:     "tens_and_ones",
			n:        25,
			expected: "twenty five",
		},
		{
			name:     "tens",
			n:        30,
			expected: "thirty",
		},
		{
			name:     "hundreds",
			n:        100,
			expected: "one hundred",
		},
		{
			name:     "hundreds_and_ones",
			n:        105,
			expected: "one hundred five",
		},
		{
			name:     "hundreds_and_tens",
			n:        110,
			expected: "one hundred ten",
		},
		{
			name:     "hundreds_and_tens",
			n:        117,
			expected: "one hundred seventeen",
		},
		{
			name:     "hundreds_tens_and_one",
			n:        125,
			expected: "one hundred twenty five",
		},
		{
			name:     "thousand",
			n:        1000,
			expected: "one thousand",
		},
		{
			name:     "thousands",
			n:        3006,
			expected: "three thousand six",
		},
		{
			name:     "thousands_tens_ones",
			n:        5066,
			expected: "five thousand sixty six",
		},
		{
			name:     "thousands_hundreds_tens_ones",
			n:        5666,
			expected: "five thousand six hundred sixty six",
		},
		{
			name:     "tens_of_thousands_thousands_hundreds_tens_ones",
			n:        55666,
			expected: "fifty five thousand six hundred sixty six",
		},
		{
			name:     "million",
			n:        1000000,
			expected: "one million",
		},
		{
			name:     "million_tens_of_thousands_thousands_hundreds_tens_ones",
			n:        1044666,
			expected: "one million forty four thousand six hundred sixty six",
		},

	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := english_int.Convert(tc.n)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
