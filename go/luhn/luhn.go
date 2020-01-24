// Package luhn is a simple checktotal formula used to validate a variety of identification numbers, such as credit card numbers and Canadian Social Insurance Numbers.
package luhn

import (
	"strings"
	"unicode"
)

// Valid checks whether card number is valid or not
func Valid(cardNum string) bool {
	// Strip spaces between digit blocks
	cardNum = strings.Replace(cardNum, " ", "", -1)

	// Strings of length 1 or less are not valid.
	if len(cardNum) < 2 {
		return false
	}

	total, odd := 0, 0

	// The first step of the Luhn algorithm is
	// to double every second digit,
	// starting from the right.
	// Depending on whether the length of
	// the card number is odd or even
	// we can find out the first position of every
	// second digit from the right
	if len(cardNum)%2 != 0 {
		odd = 1
	}

	for i, c := range cardNum {
		// All non-digit characters are disallowed.
		if !unicode.IsDigit(c) {
			return false
		}
		// In order to convert a rune to
		// an integer you will need to
		// substract 48 to *c* and
		// you will get the converted value. Why 48?
		// Well 48 is the representation of the "0"
		// in the ascii table and since all the integers
		// are sequential in the ASCII representation,
		// this will return the desired value.
		// [1](https://tutorials.technology/tutorials/go-convert-rune-to-int.html)
		digit := int(c - '0')

		// Depending on position of the digit
		// add it or its double to the total
		if i%2 == odd {
			total += doubleDigit(digit)
		} else {
			total += digit
		}
	}
	// If the total is evenly divisible by 10,
	// then the number is valid.
	if total%10 != 0 {
		return false
	}
	return true
}

func doubleDigit(digit int) int {
	product := digit * 2
	// If doubling the number results in a number
	// greater than 9 then subtract 9 from the product.
	if product > 9 {
		return product - 9
	}
	return product
}
