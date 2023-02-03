package move_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantNums []int
	}{
		{nums: []int{0, 1}, wantNums: []int{1, 0}},
		{nums: []int{0, 1, 2}, wantNums: []int{1, 2, 0}},
		{nums: []int{1}, wantNums: []int{1}},
		{nums: []int{0, 0, 0, 0}, wantNums: []int{0, 0, 0, 0}},
		{nums: []int{0, 1, 0, 3, 12}, wantNums: []int{1, 3, 12, 0, 0}},
		{nums: []int{0, 0, 1, 3, 12}, wantNums: []int{1, 3, 12, 0, 0}},
		{nums: []int{0, 0, 0, 0, 12}, wantNums: []int{12, 0, 0, 0, 0}},
		{nums: []int{0, 0, 12, 0, 0}, wantNums: []int{12, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.nums
			moveZeroes(nums)
			assert.Equal(t, tt.wantNums, nums)
		})
	}
}

func Test_moveToTheEnd(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		pos      int
		wantNums []int
	}{
		{nums: []int{0, 1, 0, 3, 12}, pos: 0, wantNums: []int{1, 0, 3, 12, 0}},
		{nums: []int{0, 1, 0, 3, 12}, pos: 3, wantNums: []int{0, 1, 0, 12, 3}},
		{nums: []int{0, 0, 12, 0, 0}, pos: 2, wantNums: []int{0, 0, 0, 0, 12}},
		{nums: []int{0, 1}, pos: 1, wantNums: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.nums
			moveToTheEnd(nums, tt.pos)
			assert.Equal(t, tt.wantNums, nums)
		})
	}
}
