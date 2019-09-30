package triangle


// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind

const (
    // Pick values for the following identifiers used by the test program.
    NaT // not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	return k
}
