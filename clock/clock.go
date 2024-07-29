package clock

import (
	"fmt"
)

type Clock struct {
	hour   int
	minute int
}

const (
	HOUR_IN_MINS int = 60
	DAY_IN_HOURS int = 24
)

func normalize(c Clock) Clock {
	h := c.hour % DAY_IN_HOURS
	if h < 0 {
		h = DAY_IN_HOURS + h
	}

	m := c.minute % HOUR_IN_MINS
	if m < 0 {
		m = HOUR_IN_MINS + m
	}

	hourDiff := int((c.minute - m) / 60) // natural digit from -Inf to +Inf
	// representing the diff that minutes create in hours

	h += int(hourDiff) // calculate the final value of the hour
	h %= DAY_IN_HOURS
	
	if h < 0 {
		return normalize(Clock{h, m})
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
	newHours := c.hour
	newMinutes := c.minute - m

	return normalize(Clock{hour: newHours, minute: newMinutes})
}

func (c Clock) String() string {
	c = normalize(c)
	h, m := c.hour, c.minute

	str := ""

	if h > 9 {
		str += fmt.Sprintf("%d", h)
	} else {
		str += fmt.Sprintf("0%d", h)
	}

	str += ":"

	if m > 9 {
		str += fmt.Sprintf("%d", m)
	} else {
		str += fmt.Sprintf("0%d", m)
	}

	return str
}
