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
