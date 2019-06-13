package smallest

import (
	"math"
	"sort"
)

func Difference(a, b []int) int {
	sort.Ints(a)
	sort.Ints(b)

	// index of a
	i := 0
	// index of b
	j := 0

	minDiff := math.MaxInt32

	for i < len(a) && j < len(b) {
		v1 := a[i]
		v2 := b[j]

		diff := Abs(v1 - v2)
		if diff < minDiff {
			minDiff = diff
		}

		if v1 < v2 {
			i++
		} else {
			j++
		}
	}

	return minDiff
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
