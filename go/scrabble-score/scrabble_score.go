package scrabble

import "strings"

// Score computes the Scrabble score for given word
func Score(word string) int {

	score := 0

	OnePointLetters := "AEIOULNRST"
	TwoPointsLetters := "DG"
	ThreePointsLetters := "BCMP"
	FourPointsLetters := "FHVWY"
	FivePointsLetters := "K"
	EightPointesLetters := "JX"
	TenPointsLetters := "QZ"

	for _, w := range strings.ToUpper(word) {
		if strings.ContainsRune(OnePointLetters, w) {
			score++
		}
		if strings.ContainsRune(TwoPointsLetters, w) {
			score += 2
		}
		if strings.ContainsRune(ThreePointsLetters, w) {
			score += 3
		}
		if strings.ContainsRune(FourPointsLetters, w) {
			score += 4
		}
		if strings.ContainsRune(FivePointsLetters, w) {
			score += 5
		}
		if strings.ContainsRune(EightPointesLetters, w) {
			score += 8
		}
		if strings.ContainsRune(TenPointsLetters, w) {
			score += 10
		}
	}

	return score
}
