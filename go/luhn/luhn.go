package luhn

import (
	"strings"
	"unicode"
)

// Valid checks whether card number is valid or not
func Valid(cardNum string) bool {
	// Remove spaces between digit blocks
	stripped := strings.Replace(cardNum, " ", "", -1)

	if len(stripped) < 2 {
		return false
	}

	sum := 0
	odd := 0
	if len(stripped)%2 != 0 {
		odd = 1
	}

	for i, c := range stripped {
		if !unicode.IsDigit(c) {
			return false
		}

		digit := int(c - '0')

		if i%2 == odd {
			sum += doubleDigit(digit)
			continue
		}

		sum += digit

	}

	if sum%10 != 0 {
		return false
	}
	return true
}

func doubleDigit(digit int) int {
	sum := digit * 2
	if sum > 9 {
		return sum - 9
	}
	return sum
}
