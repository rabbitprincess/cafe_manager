package utilx

import (
	"regexp"
)

// 01x-xxxx-xxxx
// 01x-xxx-xxxx
var (
	phoneRegex = regexp.MustCompile(`^01[016789]-\d{3,4}-\d{4}$`)
)

func IsPhoneNumber(s string) bool {
	return phoneRegex.MatchString(s)
}
