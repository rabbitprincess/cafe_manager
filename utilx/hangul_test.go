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

func TestGetInitialFromHangul(t *testing.T) {
	for _, test := range []struct {
		input  string
		expect string
	}{
		{"ㄱㄴㄷ", "ㄱㄴㄷ"},
		{"ㄲㄸㅃㅆㅉ", "ㄲㄸㅃㅆㅉ"},
		{"김철수", "ㄱㅊㅅ"},
		{"박영희", "ㅂㅇㅎ"},
		{"홍길동", "ㅎㄱㄷ"},
		{"안녕하세요세상", "ㅇㄴㅎㅅㅇㅅㅅ"},
		{"가나다라마바사아자차카타파하", "ㄱㄴㄷㄹㅁㅂㅅㅇㅈㅊㅋㅌㅍㅎ"},
		{"각난닫람맏밥삿앙잦찿캌탙팦핳", "ㄱㄴㄷㄹㅁㅂㅅㅇㅈㅊㅋㅌㅍㅎ"},
		{"까따빠싸짜", "ㄲㄸㅃㅆㅉ"},
	} {
		result := GetInitialFromHangul(test.input)
		require.Equal(t, test.expect, result, test.input)
	}
}
