package reverse

import "strings"

// Reverse input string, safe with UTF chars
func Reverse(word string) string {
	unicodedStr := []rune(word)

	var sb strings.Builder

	for i, _ := range unicodedStr {
		sb.WriteRune(unicodedStr[len(unicodedStr)-1-i])
	}
	return sb.String()
}
