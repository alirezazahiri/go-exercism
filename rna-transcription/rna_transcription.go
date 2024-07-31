package strand

func dnaToRna(r rune) rune {
	switch r {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	default:
		return -1
	}
}
func ToRNA(dna string) string {
	res := ""

	for _, ch := range dna {
		res += string(dnaToRna(ch))
	}

	return res
}
