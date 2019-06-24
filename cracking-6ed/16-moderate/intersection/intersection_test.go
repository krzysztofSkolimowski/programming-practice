package intersection_test

import (
	"testing"

	"github.com/krzysztofSkolimowski/programming-practice/cracking-6ed/16-moderate/intersection"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name             string
		beginningA, endA intersection.Point
		beginningB, endB intersection.Point
		expected         bool
	}{
		{
			name:       "same_line_vertical",
			beginningA: intersection.Point{0, 0},
			endA:       intersection.Point{0, 1},
			beginningB: intersection.Point{0, 0},
			endB:       intersection.Point{0, 2},
			expected:   false,
		},
		{

			name:       "same_line_horizontal",
			beginningA: intersection.Point{0, 0},
			endA:       intersection.Point{1, 0},
			beginningB: intersection.Point{0, 0},
			endB:       intersection.Point{2, 0},
			expected:   false,
		},
		{
			name:       "parallel_lines",
			beginningA: intersection.Point{0, 0},
			endA:       intersection.Point{2, 2},
			beginningB: intersection.Point{2, 0},
			endB:       intersection.Point{4, 2},
			expected:   false,
		},
		{
			name:       "perpendicular",
			beginningA: intersection.Point{0, 0},
			endA:       intersection.Point{1, 0},
			beginningB: intersection.Point{1, 0},
			endB:       intersection.Point{0, 2},
			expected:   false,
		},
		{
			name:       "not_intersecting",
			beginningA: intersection.Point{0, 0},
			endA:       intersection.Point{0, 5},
			beginningB: intersection.Point{2, 0},
			endB:       intersection.Point{0, 6},
			expected:   false,
		},
		{
			name:       "intersecting",
			beginningA: intersection.Point{0, 0},
			endA:       intersection.Point{4, 5},
			beginningB: intersection.Point{2, 0},
			endB:       intersection.Point{0, 6},
			expected:   true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lineA, _ := intersection.NewLine(tc.beginningA, tc.endA)
			lineB, _ := intersection.NewLine(tc.beginningB, tc.endB)


			if actual := intersection.Intersect(lineA, lineB); actual != tc.expected {
				t.Errorf("Intersect() = %v, expected %v", actual, tc.expected)
			}
		})
	}
}

func Test_line_Parallel(t *testing.T) {
	tests := []struct {
		name             string
		beginningA, endA intersection.Point
		beginningB, endB intersection.Point
		expected         bool
	}{
		{
			name:       "same_line_vertical",
			beginningA: intersection.Point{0, 0},
			endA:       intersection.Point{0, 1},
			beginningB: intersection.Point{0, 0},
			endB:       intersection.Point{0, 2},
			expected:   true,
		},
		{

			name:       "same_line_horizontal",
			beginningA: intersection.Point{0, 0},
			endA:       intersection.Point{1, 0},
			beginningB: intersection.Point{0, 0},
			endB:       intersection.Point{2, 0},
			expected:   true,
		},
		{
			name:       "parallel_lines",
			beginningA: intersection.Point{0, 0},
			endA:       intersection.Point{2, 2},
			beginningB: intersection.Point{2, 0},
			endB:       intersection.Point{4, 2},
			expected:   true,
		},
		{
			name:       "perpendicular",
			beginningA: intersection.Point{0, 0},
			endA:       intersection.Point{1, 0},
			beginningB: intersection.Point{0, 0},
			endB:       intersection.Point{0, 1},
			expected:   false,
		},
		{
			name:       "not_parallel",
			beginningA: intersection.Point{0, 0},
			endA:       intersection.Point{6, 2},
			beginningB: intersection.Point{0, 0},
			endB:       intersection.Point{1, 3},
			expected:   false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lineA, _ := intersection.NewLine(tc.beginningA, tc.endA)
			lineB, _ := intersection.NewLine(tc.beginningB, tc.endB)

			if got := lineA.Parallel(lineB); got != tc.expected {
				t.Errorf("line.Parallel() = %v, expected %v", got, tc.expected)
			}
		})
	}
}
