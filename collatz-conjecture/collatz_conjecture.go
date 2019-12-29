package collatzconjecture

import "errors"

func CollatzConjecture(in int) (int, error) {
	steps := 0

	if in < 1 {
		return 0, errors.New("Invalid input, only positive integers allowed!")
	}

	for in > 1 {
		if in&1 == 1 {
			in = (in*3 + 1)
		} else {
			in = in / 2
		}
		steps++
	}

	return steps, nil
}
