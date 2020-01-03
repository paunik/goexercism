package clock

import "fmt"

// Clock represents clock implementation
type Clock struct {
	m int
}

// New constructor for Clock Type
func New(h int, m int) Clock {
	h, m = hourAndMinute(h*60 + m)
	return Clock{h*60 + m}
}

// Add method adds minutes to current instance
func (c Clock) Add(a int) Clock {
	h, m := hourAndMinute(c.m + a)
	c.m = h*60 + m
	return c
}

// Subtract method subtract minutes to current instance
func (c Clock) Subtract(s int) Clock {
	h, m := hourAndMinute(c.m - s)
	c.m = h*60 + m
	return c
}

// String stringer interface implementation
func (c Clock) String() string {
	h, m := hourAndMinute(c.m)
	return fmt.Sprintf("%0.2d:%0.2d", h, m)
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
