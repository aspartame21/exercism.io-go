package grains

import (
	"errors"
)

// Square returns the amount of grains in a single square
func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, errors.New("out of range number, expected between 1 and 64 inclusively")
	}
	return 1 << (num - 1), nil
}

// Total calculates the amount of grains in all 64 squares
func Total() uint64 {
	return 1<<64 - 1
}
