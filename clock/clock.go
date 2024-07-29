package clock

import (
	"fmt"
)

type Clock struct {
	hour   int
	minute int
}

const (
	hourInMinutes int = 60
	dayInHours    int = 24
)

func normalize(c Clock) Clock {
	h := (c.hour + (c.minute / hourInMinutes)) % dayInHours
	m := c.minute % hourInMinutes

	if m < 0 {
		m += hourInMinutes
		h -= 1
	}
	if h < 0 {
		h += dayInHours
	}

	return Clock{h, m}
}

func New(h, m int) Clock {
	return normalize(Clock{h, m})
}

func (c1 Clock) Equal(c2 Clock) bool {
	nc1, nc2 := normalize(c1), normalize(c2)
	return nc1.hour == nc2.hour && nc1.minute == nc2.minute
}

func (c Clock) Add(m int) Clock {
	newHours := c.hour
	newMinutes := c.minute + m

	return normalize(Clock{hour: newHours, minute: newMinutes})
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
