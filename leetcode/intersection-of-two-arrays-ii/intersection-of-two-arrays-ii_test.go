package intersection_of_two_arrays_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersect(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, want: []int{2, 2}},
		{nums1: []int{1, 1, 1, 1}, nums2: []int{2, 2}, want: []int{}},
		{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, want: []int{4, 9}},
		{nums1: []int{1, 1, 1, 1}, nums2: []int{1, 1, 1, 1}, want: []int{1, 1, 1, 1}},
		{nums1: []int{1}, nums2: []int{1, 1, 1, 1}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, intersect(tt.nums1, tt.nums2))
		})
	}
}
