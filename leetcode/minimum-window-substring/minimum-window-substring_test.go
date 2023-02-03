package minimum_window_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{s: "ADOBECODEBANC", t: "ABC", want: "BANC"},
		{s: "ABC", t: "ABC", want: "ABC"},
		{s: "ABCABC", t: "ABC", want: "ABC"},
		{s: "ASSSBSSSC", t: "ABC", want: "ASSSBSSSC"},
		{s: "a", t: "a", want: "a"},
		{s: "a", t: "aa", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minWindow(tt.s, tt.t))
		})
	}
}
