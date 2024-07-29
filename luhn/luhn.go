package luhn

import "strings"

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 {
		return false
	}
	numId := []rune(id) // convert stripped id string to a list of ASCII codes
	isSecond := false
	sum := 0
	for i := len(numId) - 1; i >= 0; i-- {
		numId[i] -= 48                    // convert ASCII code to digit (48 is the ASCII code of '0')
		if numId[i] > 9 || numId[i] < 0 { // if it is non digit
			return false
		}
		if isSecond {
			doubled := numId[i] * 2
			if doubled > 9 {
				doubled -= 9
			}
			numId[i] = doubled
		}
		sum += int(numId[i])
		isSecond = !isSecond
	}
	return sum%10 == 0
}
