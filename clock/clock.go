package clock

import "fmt"

//TestVersion checks if this code is running on the right test.
const TestVersion = 2

//Clock represents a hour and minute withou a date
type Clock struct {
	Hour, Minute int
}

//Time returns a clock with the corresponding hour and minutes
func Time(hour, minute int) Clock {
	extraHour := minute / 60
	return Clock{hour%24 + extraHour, minute % 60}
}

//Add minutes to this clock
func (c Clock) Add(minutes int) Clock {
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}
