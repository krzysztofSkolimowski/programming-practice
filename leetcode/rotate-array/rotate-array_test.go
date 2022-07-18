package rotate_array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		expectedNums []int
		k            int
	}{
		{k: 0, nums: []int{1, 2, 3, 4, 5, 6, 7}, expectedNums: []int{1, 2, 3, 4, 5, 6, 7}},
		{k: 2, nums: []int{1, 2, 3, 4, 5, 6, 7}, expectedNums: []int{6, 7, 1, 2, 3, 4, 5}},
		{k: 2, nums: []int{1, 2, 3, 4, 5, 6, 7}, expectedNums: []int{6, 7, 1, 2, 3, 4, 5}},
		{k: 3, nums: []int{1, 2, 3, 4, 5, 6, 7}, expectedNums: []int{5, 6, 7, 1, 2, 3, 4}},
		{k: 7, nums: []int{1, 2, 3, 4, 5, 6, 7}, expectedNums: []int{1, 2, 3, 4, 5, 6, 7}},
		{k: 7, nums: []int{1, 2, 3}, expectedNums: []int{3, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.nums
			rotate(nums, tt.k)
			fmt.Println(nums)
			assert.Equal(t, tt.expectedNums, nums)
		})
	}
}
