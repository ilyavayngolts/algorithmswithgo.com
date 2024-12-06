package module01

import (
	"strings"
	"unicode/utf8"
)

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	var reversedStringBuilder strings.Builder
	for i, w := len(word), 0; i > 0; i -= w {
		var r rune
		r, w = utf8.DecodeLastRuneInString(word[:i])
		reversedStringBuilder.WriteRune(r)
	}

	return reversedStringBuilder.String()
}
