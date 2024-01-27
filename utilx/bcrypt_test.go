package utilx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestEncryptCheck(t *testing.T) {
	for _, test := range []struct {
		name            string
		inputPassword   string
		comparePassword string
		expectError     error
		check           bool
	}{
		{"success case 1", "hello world!", "hello world!", nil, true},
		{"success case 2", "123456789!@#$%^&*()", "123456789!@#$%^&*()", nil, true},
		{"success case 3", "한국어비밀번호", "한국어비밀번호", nil, true},
		{"success case 4", "1234", "1234", nil, true},
		{"success case 4", strings.Repeat("w", 72), strings.Repeat("w", 72), nil, true},
		{"fail case 1", "hello world!", "hello wordl!", nil, false},
		{"except case 1", strings.Repeat("w", 72) + "a", strings.Repeat("w", 72) + "a", bcrypt.ErrPasswordTooLong, false},
	} {
		hash, err := BEncrypt(test.inputPassword)
		if err == test.expectError {
			continue
		}
		require.NoError(t, err, test.name)

		check, err := BCheck(test.comparePassword, hash)
		if err == test.expectError {
			continue
		}
		require.NoError(t, err, test.name)

		require.Equal(t, test.check, check, test.name)
	}
}
