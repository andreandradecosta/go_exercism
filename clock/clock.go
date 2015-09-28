package clock

import "fmt"

//TestVersion checks if this code is running on the right test.
const TestVersion = 2

//Clock represents a hour and minute withou a date
type Clock struct {
	Hour, Minute int
}

func (m modulo) addModulo(mod int, a int) int {

}

//Time returns a clock with the corresponding hour and minutes
func Time(hour, minute int) Clock {
	hour = hour % 24
	extraHour := minute / 60
	hour = hour + extraHour
	if hour < 0 {
		hour = 24 + hour
	}
	return Clock{hour, minute % 60}
}

//Add minutes to this clock
func (c Clock) Add(minutes int) Clock {
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}
