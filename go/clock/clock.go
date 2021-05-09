package clock

import (
	"fmt"
)

const minutesInAnHour = 60
const minutesInADay = 24 * minutesInAnHour

type Clock struct {
	minutes int
}

func New(h, m int) Clock {
	minutes := (h*minutesInAnHour + m) % minutesInADay
	if minutes < 0 {
		minutes += minutesInADay
	}
	return Clock{minutes: minutes}
}

func (c Clock) Add(m int) Clock {
	return New(0, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}
