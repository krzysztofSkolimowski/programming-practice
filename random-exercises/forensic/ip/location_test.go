package ip

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func Test_locationFetcher_Country(t *testing.T) {
	tests := []struct {
		name          string
		ip            net.IP
		expected      string
		expectedError error
	}{
		{
			name:          "brazil",
			ip:            net.ParseIP("187.12.249.74"),
			expected:      "Brazil",
			expectedError: nil,
		},
		{
			name:          "brazil",
			ip:            net.ParseIP("187.12.249.74"),
			expected:      "Brazil",
			expectedError: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f, err := NewLocationFetcher()
			assert.NoError(t, err)

			actual, err := f.Country(tc.ip)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
