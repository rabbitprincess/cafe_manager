package utilx

import (
	"unicode"
	"unicode/utf8"
)

func IsHangulInitial(ch rune) bool {
	// 초성에 해당하는 Unicode 코드 포인트 범위
	const (
		startRune = 0x1100
		endRune   = 0x115f
	)

	return startRune <= ch && ch <= endRune
}

func IsHangulInitialsOnly(s string) bool {
	for _, ch := range s {
		// 초성만으로 구성되었는지 확인
		_, size := utf8.DecodeRuneInString(string(ch))
		if size > 1 && !IsHangulInitial(ch) {
			return false
		}
	}
	return true
}

func IsHangulOnly(s string) bool {
	for _, ch := range s {
		if !unicode.Is(unicode.Hangul, ch) {
			return false
		}
	}
	return true
}
