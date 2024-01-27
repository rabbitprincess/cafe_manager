package utilx

import "unicode"

func IsHangulInitial(ch rune) bool {
	const (
		startRune = 0x3131 // 'ㄱ'
		endRune   = 0x314E // 'ㅎ'
	)

	return startRune <= ch && ch <= endRune
}

func IsHangulInitialsOnly(s string) bool {
	for _, ch := range s {
		if !IsHangulInitial(ch) {
			return false
		}
	}
	return true
}

func IsHangul(ch rune) bool {
	return unicode.Is(unicode.Hangul, ch)
}

func IsHangulOnly(s string) bool {
	for _, ch := range s {
		if !IsHangul(ch) {
			return false
		}
	}
	return true
}

func GetInitialFromHangul(s string) string {
	var result string
	for _, ch := range s {
		if IsHangulInitial(ch) {
			result += string((ch))
		} else if IsHangul(ch) {
			result += string(convertChosungToLetter(((ch - 0xAC00) / (28 * 21)) + 0x1100))
		}
	}
	return result
}

func convertChosungToLetter(ch rune) rune {
	switch ch {
	case 'ᄀ':
		return 'ㄱ'
	case 'ᄁ':
		return 'ㄲ'
	case 'ᄂ':
		return 'ㄴ'
	case 'ᄃ':
		return 'ㄷ'
	case 'ᄄ':
		return 'ㄸ'
	case 'ᄅ':
		return 'ㄹ'
	case 'ᄆ':
		return 'ㅁ'
	case 'ᄇ':
		return 'ㅂ'
	case 'ᄈ':
		return 'ㅃ'
	case 'ᄉ':
		return 'ㅅ'
	case 'ᄊ':
		return 'ㅆ'
	case 'ᄋ':
		return 'ㅇ'
	case 'ᄌ':
		return 'ㅈ'
	case 'ᄍ':
		return 'ㅉ'
	case 'ᄎ':
		return 'ㅊ'
	case 'ᄏ':
		return 'ㅋ'
	case 'ᄐ':
		return 'ㅌ'
	case 'ᄑ':
		return 'ㅍ'
	case 'ᄒ':
		return 'ㅎ'
	default:
		return ch
	}
}
