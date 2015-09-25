package leap

//Checks the version of the test
const TestVersion = 1

//IsLeapYear returns true if its a leap year
func IsLeapYear(year int) bool {
	switch {
	case year%4 > 0:
		return false
	case year%100 > 0:
		return true
	case year%400 > 0:
		return false
	}
	return true
}
