package utilx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsHangul(t *testing.T) {
	for _, test := range []struct {
		input    string
		isHangul bool
	}{
		{"가", true},
		{"박", true},
		{"ㄱ", true},
		{"ㅏ", true},
		{"꿿", true},
		{"찦", true},
		{"쑛", true},
		{"똠", true},
		{"a", false},
		{"A", false},
		{"1", false},
		{"!", false},
		{"あ", false},
		{"我", false},
		{"é", false},
		{"ñ", false},
		{"안녕하세요세상", true},
		{"안녕하세요세상!", false},
	} {
		require.Equal(t, test.isHangul, IsHangulOnly(test.input), test.input)
	}
}

func TestIsHangulInitial(t *testing.T) {
	for _, test := range []struct {
		input           string
		isHangulInitial bool
	}{
		{"ㄱ", true},
		{"ㄴ", true},
		{"ㅎ", true},
		{"ㅏ", false},
		{"ㅣ", false},
		{"가", false},
		{"박", false},
		{"a", false},
		{"A", false},
		{"1", false},
		{"!", false},
		{"あ", false},
		{"我", false},
		{"é", false},
		{"ñ", false},
		{"ㅁㄴㅇㄹ", true},
		{"ㅁㄴㅇㄹㅏㅑㅓㅕ", false},
	} {
		require.Equal(t, test.isHangulInitial, IsHangulInitialsOnly(test.input), test.input)
	}
}

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
