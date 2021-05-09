package clock

import (
	"fmt"
)

type Clock struct {
	minutes int
}

func New(h, m int) Clock {
	computed := normalize(h*60 + m)
	return Clock{minutes: computed}
}

func (c Clock) Add(n int) Clock {
	c.minutes = normalize(c.minutes + n)
	return c
}

func (c Clock) Subtract(n int) Clock {
	c.minutes = normalize(c.minutes - n)
	return c
}

func (c Clock) String() string {
	c.minutes = normalize(c.minutes)
	return fmt.Sprintf("%02d:%02d", c.minutes/60%24, c.minutes%60)
}

func normalize(minutes int) int {
	minutes %= 1440
	if minutes < 0 {
		minutes += 1440
	}
	return minutes
}
