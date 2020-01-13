package isogram

import "strings"

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for i, letter := range word {
		source := word[0:i] + word[i+1:len(word)]

		if letter == '-' || letter == ' ' {
			continue
		}

		if strings.ContainsRune(source, letter) {
			return false
		}
	}

	return true
}
