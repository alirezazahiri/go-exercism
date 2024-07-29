package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("inputs should be of equal length")
	}

	diff := 0
	runedB := []rune(b)
	for i, ch := range a {
		if runedB[i] != ch {
			diff++
		}
	}

	return diff, nil
}
