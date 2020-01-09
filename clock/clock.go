package clock

import "fmt"

const MinutesInDay = 24 * 60
const MinutesInHour = 60

// Clock is clock that is very efficient as don't use hours.
type Clock int

// New creates a new clock.
func New(h, m int) Clock {
	minutes := (h*MinutesInHour + m) % MinutesInDay
	if minutes < 0 {
		minutes += MinutesInDay
	}
	return Clock(minutes)
}

// String returns a string representation of the clock.
func (c Clock) String() string {
	h := c / MinutesInHour
	return fmt.Sprintf("%02d:%02d", h, c-h*MinutesInHour)
}

// Add creates a new clock that is `minute` ahead from this one.
func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

// Subtract method subtract minutes to current instance
func (c Clock) Subtract(s int) Clock {
	return New(0, int(c)-s)
}
