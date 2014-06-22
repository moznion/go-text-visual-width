package textvidualwidth

import "github.com/moznion/go-unicode-east-asian-width"

func Width(str string) int {
	var length int
	for _, char := range str {
		length++
		if eastasianwidth.IsFullwidth(char) {
			length++
		}
	}

	return length
}

func Trim(str string, limit int) string {
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