package raindrops

import "strconv"

// Convert converts a number to a string,
// the contents of which depend
// on the number's factors.
func Convert(num int) string {

	var raindrop string

	if num%3 == 0 {
		raindrop += "Pling"
	}
	if num%5 == 0 {
		raindrop += "Plang"
	}
	if num%7 == 0 {
		raindrop += "Plong"
	}
	if raindrop == "" {
		raindrop = strconv.Itoa(num)
	}

	return raindrop
}
