package scrabble

import "strings"

type lettersPointsMap struct {
	letters string
	points  int
}

// Score computes the Scrabble score for given word
func Score(word string) int {

	score := 0
	lpm := []lettersPointsMap{
		{"AEIOULNRST", 1},
		{"DG", 2},
		{"BCMP", 3},
		{"FHVWY", 4},
		{"K", 5},
		{"JX", 8},
		{"QZ", 10},
	}

	for _, w := range strings.ToUpper(word) {
		for _, mapping := range lpm {
			if strings.ContainsRune(mapping.letters, w) {
				score += mapping.points
			}
		}
	}

	return score
}
