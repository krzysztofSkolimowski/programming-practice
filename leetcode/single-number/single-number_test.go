package single_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{want: 1, nums: []int{2, 2, 1}},
		{want: 4, nums: []int{4, 1, 2, 1, 2}},
		{want: 3, nums: []int{3, 1, 2, 1, 2, 6, 6, 7, 7}},
		{want: 1, nums: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, singleNumber(tt.nums))
		})
	}
}
