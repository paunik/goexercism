package clock

import "fmt"

// Clock represents clock implementation
type Clock struct {
	h int
	m int
}

// New constructor for Clock Type
func New(h int, m int) Clock {
	min := h*60 + m
	h, m = hourAndMinute(min)

	return Clock{h, m}
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

	c.h, c.m = hourAndMinute(min)

	return c
}

// String stringer interface implementation
func (c Clock) String() string {
	return fmt.Sprintf("%0.2d:%0.2d", c.h, c.m)
}

// hourAndMinute calculate hour and minute based on total number of mintues
func hourAndMinute(min int) (h, m int) {
	if min >= 0 {
		h = (min / 60) % 24
		m = min % 60
	} else {
		min = -min
		h = 23 - (min/60)%24
		m = 60 - (min % 60)
	}

	if h == 24 {
		h = 0
	}

	if m == 60 {
		m = 0
		h = (h + 1) % 24
	}
	return
}
