package best_time_to_buy_and_sell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		//{
		//	prices: []int{7, 6, 4, 3, 1},
		//	want:   0,
		//},
		//{
		//	prices: []int{1, 5, 3, 2, 5, 7},
		//	want:   9,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxProfit(tt.prices))
		})
	}
}
