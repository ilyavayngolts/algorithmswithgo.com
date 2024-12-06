package module01

import "strings"

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	var reversedStringBuilder strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		reversedStringBuilder.WriteByte(word[i])
	}

	return reversedStringBuilder.String()
}
