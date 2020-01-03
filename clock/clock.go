package clock

import "fmt"

const HoursInDay = 24
const MinutesInHour = 60
const LastHour = 23

// Clock represents clock implementation
type Clock struct {
	m int
}

// New constructor for Clock Type
func New(h int, m int) Clock {
	h, m = hourAndMinute(h*MinutesInHour + m)
	return Clock{h*MinutesInHour + m}
}

// Add method adds minutes to current instance
func (c Clock) Add(a int) Clock {
	h, m := hourAndMinute(c.m + a)
	c.m = h*MinutesInHour + m
	return c
}

// Subtract method subtract minutes to current instance
func (c Clock) Subtract(s int) Clock {
	h, m := hourAndMinute(c.m - s)
	c.m = h*MinutesInHour + m
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
		h = (min / MinutesInHour) % HoursInDay
		m = min % MinutesInHour
	} else {
		min = -min
		h = LastHour - (min/MinutesInHour)%HoursInDay
		m = MinutesInHour - (min % MinutesInHour)
	}

	if h == HoursInDay {
		h = 0
	}

	if m == MinutesInHour {
		m = 0
		h = (h + 1) % HoursInDay
	}

	return
}
