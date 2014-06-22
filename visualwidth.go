package visualwidth

import "github.com/moznion/go-unicode-east-asian-width"

// EastAsian is the flag for ambiguous character.
// If this flag is `true`, ambiguous characters are treated as full width.
// Elsewise, ambiguous characters are treated as half width.
// Default value is `false`.
var EastAsian = false

// Width returns visual width of string.
// This function treats the length of full width character is 2, other hands is 1.
func Width(str string) int {
	eastasianwidth.EastAsian = EastAsian

	var length int
	for _, char := range str {
		length++
		if eastasianwidth.IsFullwidth(char) {
			length++
		}
	}

	return length
}

// Trim returns the string that is trimmed by specified limit length.
func Trim(str string, limit int) string {
	eastasianwidth.EastAsian = EastAsian

	var count int
	var ret string
	for _, char := range str {
		if eastasianwidth.IsFullwidth(char) {
			count += 2
			if count > limit {
				break
			}
			ret += string(char)
		} else {
			count++
			if count > limit {
				break
			}
			ret += string(char)
		}
	}

	return ret
}
