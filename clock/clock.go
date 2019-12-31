package clock

import "fmt"

// Clock represents clock implementation
type Clock struct {
	h int
	m int
}

// New constructor for Clock Type
func New(h int, m int) Clock {
	c := Clock{0, 0}

	min := h*60 + m
	positive := min >= 0

	if positive {
		c.h = (min / 60) % 24
		c.m = min % 60
	} else {
		min = -min
		c.h = 23 - (min/60)%24
		c.m = 60 - (min % 60)
	}

	if c.h == 24 {
		c.h = 0
	}

	if c.m == 60 {
		c.m = 0
		c.h = (c.h + 1) % 24
	}

	return c
}

// Add method adds minutes to current instance
func (c Clock) Add(a int) Clock {
	t := c.h*60 + c.m + a
	c.h = (t / 60) % 24
	c.m = t % 60

	return c
}

// Subtract method subtract minutes to current instance
func (c Clock) Subtract(s int) Clock {
	min := c.h*60 + c.m - s
	positive := min >= 0

	if positive {
		c.h = (min / 60) % 24
		c.m = min % 60
	} else {
		min = -min
		c.h = 23 - (min/60)%24
		c.m = 60 - (min % 60)
	}

	if c.h == 24 {
		c.h = 0
	}

	if c.m == 60 {
		c.m = 0
		c.h = (c.h + 1) % 24
	}

	return c
}

// String stringer interface implementation
func (c Clock) String() string {
	return fmt.Sprintf("%0.2d:%0.2d", c.h, c.m)
}
