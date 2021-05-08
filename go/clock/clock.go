package clock

import (
	"fmt"
	"strconv"
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
	return fmt.Sprintf("%s:%s", format((c.minutes/60)%24), format(c.minutes%60))
}

func normalize(minutes int) int {
	minutes %= 1440
	if minutes < 0 {
		minutes += 1440
	}
	return minutes
}

func format(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}
