// Package clock provides Clock creation, Add/Subtract, and String representation methods.
package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour, minute int
}

// New creates a new Clock
func New(hour, minute int) Clock {
	return Clock{hour, minute}.format()
}

// String returns a string representation of a Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds a positive/negative number of minutes to a Clock
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return c.format()
}

func (c Clock) format() Clock {
	for c.minute < 0 {
		c.minute += 60
		c.hour -= 1
	}
	for c.hour < 0 {
		c.hour += 24
	}

	c.hour += c.minute / 60
	c.minute = c.minute % 60

	c.hour = c.hour % 24
	return c
}
