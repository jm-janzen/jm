package word

import "strings"

// Returns string with the first letter capitalised
// That's right, it won't work on strings with multiple words,
// which is why this package is called word and not words, duh. 😛️
func Capitalise(s string) string {
	return strings.ToTitle(string(s[0])) + s[1:]
}
