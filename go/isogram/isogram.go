package isogram

import "strings"

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for i, letter := range word {

		if letter == '-' || letter == ' ' {
			continue
		}

		index := strings.IndexRune(word, letter)

		if index != i && index != -1 {
			return false
		}
	}

	return true
}
