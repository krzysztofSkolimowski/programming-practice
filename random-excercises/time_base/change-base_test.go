package time_base_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeTime(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			n:    0,
			want: "",
		},
		{
			n:    50,
			want: " 50 sec",
		},
		{
			n:    60,
			want: " 1 min",
		},
		{
			n:    3 * 60,
			want: " 3 min",
		},
		{
			n:    1*60 + 10,
			want: " 1 min 10 sec",
		},
		{
			n:    1 * 60 * 60,
			want: " 1 h",
		},
		{
			n:    60*60*5 + 6,
			want: " 5 h 6 sec",
		},
		{
			n:    60*60*5 + 6*60,
			want: " 5 h 6 min",
		},
		{
			n:    60*60*5 + 5*60 + 6,
			want: " 5 h 5 min 6 sec",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ChangeTime(tc.n)
			assert.Equal(t, tc.want, result)
		})
	}
}
