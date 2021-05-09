package romannumerals

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var RomanNumeralsTable = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

var InputError = errors.New("input must be between 1 and 3000")

func ToRomanNumeral(arabic int) (string, error) {

	if arabic < 1 || arabic > 3000 {
		return "", InputError
	}

	var result strings.Builder

	for _, numeral := range RomanNumeralsTable {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String(), nil
}
