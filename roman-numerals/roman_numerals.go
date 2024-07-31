package romannumerals

import "errors"

type Int int
type String string

func getRomanNumber(roman_number string) Int {
	switch roman_number {
	case "I":
		return 1
	case "IV":
		return 4
	case "V":
		return 5
	case "IX":
		return 9
	case "X":
		return 10
	case "XL":
		return 40
	case "L":
		return 50
	case "XC":
		return 90
	case "C":
		return 100
	case "CD":
		return 400
	case "D":
		return 500
	case "CM":
		return 900
	case "M":
		return 1000
	default:
		return 0
	}
}

func (s *String) append(a String, repeat Int) {
	for i := 0; i < int(repeat); i++ {
		*s += a
	}
}

func (num *Int) calc(roman_number string) String {
	rn := getRomanNumber(roman_number)
	res := String("")
	if *num == 0 {
		return ""
	}
	for *num >= rn {
		count := Int((*num - (*num % rn)) / rn)
		*num -= rn * count
		res.append(String(roman_number), count)
	}
	return res
}

func ToRomanNumeral(input int) (string, error) {
	if input >= 4000 || input <= 0 {
		return "", errors.New("out of range")
	}
	var romanNumber Int = Int(input)
	res := String("")

	res += romanNumber.calc("M")
	res += romanNumber.calc("CM")
	res += romanNumber.calc("D")
	res += romanNumber.calc("CD")
	res += romanNumber.calc("C")
	res += romanNumber.calc("XC")
	res += romanNumber.calc("L")
	res += romanNumber.calc("XL")
	res += romanNumber.calc("X")
	res += romanNumber.calc("IX")
	res += romanNumber.calc("V")
	res += romanNumber.calc("IV")
	res += romanNumber.calc("I")

	return string(res), nil
}
