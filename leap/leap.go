package leap

func IsLeapYear(year int) bool {
	if year%4 == 0 {
		switch {
		case year%100 == 0:
			return year%400 == 0
		default:
			return true
		}
	}
	return false
}
