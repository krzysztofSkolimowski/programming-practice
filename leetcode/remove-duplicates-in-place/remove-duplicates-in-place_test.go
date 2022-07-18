package remove_duplicates_in_place

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int

		wantNums []int
		wantK    int
	}{
		{
			nums:     []int{1, 1, 2},
			wantNums: []int{1, 2, 0},
			wantK:    2,
		},
		{
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			wantNums: []int{0, 1, 2, 3, 4, 0, 0, 0, 0, 0},
			wantK:    5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.nums
			k := removeDuplicates(nums)

			assert.Equal(t, tt.wantNums, nums)
			assert.Equal(t, tt.wantK, k)
		})
	}
}
