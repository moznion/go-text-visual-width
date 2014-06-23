[![Build Status](https://travis-ci.org/moznion/go-text-visual-width.svg?branch=master)](https://travis-ci.org/moznion/go-text-visual-width)

[https://godoc.org/github.com/moznion/go-text-visual-width](https://godoc.org/github.com/moznion/go-text-visual-width)

# go-text-visual-width

Trims text by the number of the columns of terminals and mobile phones.  
This package is port of [Text::VisualWidth](https://metacpan.org/pod/Text::VisualWidth) and [Text::VisualWidth::PP](https://metacpan.org/pod/Text::VisualWidth::PP) from Perl to Go.

## Getting Started

```
go get github.com/moznion/go-text-visual-width
```

## Synopsis

```go
import "github.com/moznion/go-text-visual-width"

func main() {
	visualwidth.Width("123abcあいうｱｲｳ") // 15
	visualwidth.Trim("123ｱｲｳあいう", 8) // 123ｱｲｳあ
}
```

## Variables

- `EastAsian bool`

EastAsian is the flag for ambiguous character.  
If this flag is `true`, ambiguous characters are treated as full width.
Elsewise, ambiguous characters are treated as half width.  
Default value is `false`.

## Functions

- `func Width(str string) int`

Returns visual width of string. This function treats the length of full width character is 2, other hands 1.

- `func Trim(str string, limit int)`

Returns the string that is trimmed by specified limit length.

- `func Separate(str string, lengthFromBeginning int) [2]string`

Separate returns the strings that has two elements as `(pre, post)`.
This function separates the string by length from the beginning.

## See Also

- [Text::VisualWidth](https://metacpan.org/pod/Text::VisualWidth)
- [Text::VisualWidth::PP](https://metacpan.org/pod/Text::VisualWidth::PP)

## LICENSE

MIT
