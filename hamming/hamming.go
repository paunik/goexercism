/*
Package hamming holds methods for calculating Hamming difference between two DNA strands.
*/
package hamming

import "errors"

// Distance method determines Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {

	ar, br := []rune(a), []rune(b)

	if len(a) != len(b) {
		return 0, errors.New("strands of equal length required")
	}

	hammingDistance := 0

	for i := range ar {
		if ar[i] != br[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
