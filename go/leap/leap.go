package leap

// IsLeapYear validates whether the year passed in to the arguments is leap or not
func IsLeapYear(y int) bool {

	if y%4 == 0 {
		if y%100 != 0 {
			return true
		}
		if y%400 == 0 {
			return true
		}
	}
	return false
}
