package utilx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPhoneNumber(t *testing.T) {
	for _, test := range []struct {
		input         string
		isPhoneNumber bool
	}{
		{"010-1234-5678", true},
		{"010-123-4567", true},
		{"010-1234-567", false},

		{"011-1234-5678", true},
		{"011-123-4567", true},

		{"019-1234-5678", true},
		{"019-123-4567", true},

		{"015-1234-5678", false},
		{"02-1234-5678", false},
	} {
		require.Equal(t, test.isPhoneNumber, IsPhoneNumber(test.input), test.input)
	}
}
