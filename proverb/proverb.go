/*
Package proverb does small rhyme for input of string
*/
package proverb

import (
	"fmt"
)

const (
	lastOne = "And all for the want of a %s."
	one     = "For want of a %s the %s was lost."
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	length := len(rhyme)

	if length == 0 {
		return []string{}
	}

	a := []string{}

	for i, p := range rhyme {
		if length == 1 {
			result := fmt.Sprintf(lastOne, p)
			return append(a, result)
		}

		if len(rhyme) == i+1 {
			a = append(a, fmt.Sprintf(lastOne, rhyme[0]))
		} else {
			a = append(a, fmt.Sprintf(one, rhyme[i], rhyme[i+1]))
		}
	}

	return a
}
