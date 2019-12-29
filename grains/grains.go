package grains

import "errors"

// Square calculates number of grains on specified field on Chess board
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("Invalid input")
	}

	return 1 << (n - 1), nil
}

// Total calculates total number of grains that we got in Chess board
func Total() uint64 {
	var s uint64
	for i := 0; i < 65; i++ {
		r, _ := Square(i)
		s += r
	}

	return s
}
