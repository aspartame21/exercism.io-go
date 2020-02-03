package grains

import (
	"errors"
	"math"
)

func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, errors.New("out of range number, expected between 1 and 64 inclusively")
	}
	return uint64(math.Pow(2, float64(num-1))), nil
}

func Total() uint64 {
	var sum uint64 = 0
	for i := 1; i < 65; i++ {
		square, _ := Square(i)
		sum += square
	}
	return sum
}
