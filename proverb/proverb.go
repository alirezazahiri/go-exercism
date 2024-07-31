// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	res := make([]string, len(rhyme))
	if len(rhyme) == 0 {
		return res
	}

	createSentence := func(a, b string) string {
		return fmt.Sprintf("For want of a %s the %s was lost.", a, b)
	}
	lastSentence := fmt.Sprintf("And all for the want of a %s.", rhyme[0])

	rhymesLength := len(rhyme)

	for i, word := range rhyme {
		if i == rhymesLength-1 {
			res[i] = lastSentence
		} else {
			res[i] = createSentence(word, rhyme[i+1])
		}
	}

	return res
}
