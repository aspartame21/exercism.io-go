package main

import (
	_ "flag"
	"fmt"

	"exercism.io/go/space-age"
)

func main() {
	// a := []rune("acdefgh")
	// b := "aCdefgh"
	// fmt.Printf("%s", a[0] == b[0])
	fmt.Println(space.Age(788940000, "Earth"))
	fmt.Println(space.Age(788940000, "Mars"))
	fmt.Println(space.Age(788940000, "Venus"))
	fmt.Println(space.Age(788940000, "Jupiter"))
	fmt.Println(space.Age(788940000, "Saturn"))
	fmt.Println(2000 % 400)
}
