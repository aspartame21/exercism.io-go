package grains

import (
	"errors"
	"math"
)

// Square returns the amount of grains in a single square
func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, errors.New("out of range number, expected between 1 and 64 inclusively")
	}
	return uint64(math.Pow(2, float64(num-1))), nil
}

// Total calculates the amount of grains in all 64 squares
func Total() uint64 {
	total, _ := Square(65)
	return total - 1
}
