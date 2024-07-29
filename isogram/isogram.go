package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	visited := make(map[rune]uint8)
	for _, ch := range strings.ToLower(word) {
		if ch == ' ' || ch == '-' {
			continue
		}
		if visited[ch]++; visited[ch] >= 2 {
			return false
		}
	}
	return true
}
