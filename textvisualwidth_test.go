package textvidualwidth

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestWidth(t *testing.T) {
	RegisterTestingT(t)
	EastAsian = false

	Expect(Width("123abcあいうｱｲｳ")).To(Equal(15))
	Expect(Width("0")).To(Equal(1))
	Expect(Width("")).To(Equal(0))
}

func TestTrim(t *testing.T) {
	RegisterTestingT(t)
	EastAsian = false

	Expect(Trim("123ｱｲｳあいう", 8)).To(Equal("123ｱｲｳあ"))
	Expect(Trim("0", 8)).To(Equal("0"))
	Expect(Trim("", 8)).To(Equal(""))
}

func TestWidthForAmbiguousCharacter(t *testing.T) {
	RegisterTestingT(t)
	EastAsian = false

	ambiguousChar := "‐"

	Expect(Width(ambiguousChar)).To(Equal(1))

	EastAsian = true
	Expect(Width(ambiguousChar)).To(Equal(2))
}

func TestTrimForAmbiguousCharacter(t *testing.T) {
	RegisterTestingT(t)
	EastAsian = false

	stringThatContainsAmbiguous := "‐‐‐"

	Expect(Trim(stringThatContainsAmbiguous, 3)).To(Equal("‐‐‐"))

	EastAsian = true
	Expect(Trim(stringThatContainsAmbiguous, 3)).To(Equal("‐"))
}
