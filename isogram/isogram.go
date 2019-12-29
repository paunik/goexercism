package isogram

import "unicode"

// IsIsogram checks if word is isogram
func IsIsogram(word string) bool {
	if len(word) == 0 {
		return true
	}
	var m = make(map[rune]int)

	for _, char := range word {
		if unicode.IsLetter(char) {
			normalizedChar := unicode.ToLower(char)
			if _, found := m[normalizedChar]; found {
				return false
			}
			m[normalizedChar]++
		}
	}

	return true
}
