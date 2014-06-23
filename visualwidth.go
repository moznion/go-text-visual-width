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
	return slice(str, limit)[0]
}

// Separate returns the array of string that has two elements as `[2]string{pre, post}`.
// This function separates the string by length from the beginning.
func Separate(str string, lengthFromBeginning int) [2]string {
	return slice(str, lengthFromBeginning)
}

func slice(str string, lengthFromBeginning int) [2]string {
	eastasianwidth.EastAsian = EastAsian

	var count int
	var pre, post string
	for index, char := range str {
		if eastasianwidth.IsFullwidth(char) {
			count += 2
			if count > lengthFromBeginning {
				post = str[index:]
				break;
			}
			pre += string(char)
		} else {
			count++
			if count > lengthFromBeginning {
				post = str[index:]
			}
			pre += string(char)
		}
	}

	return [2]string{pre, post}
}
