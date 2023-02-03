package plus_one

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{digits: []int{1}, want: []int{2}},
		{digits: []int{9}, want: []int{1, 0}},
		{digits: []int{1, 0}, want: []int{1, 1}},
		{digits: []int{1, 9}, want: []int{2, 0}},
		{digits: []int{9, 9}, want: []int{1, 0, 0}},
		{digits: []int{4, 3, 2, 1}, want: []int{4, 3, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, plusOne(tt.digits))
		})
	}
}
