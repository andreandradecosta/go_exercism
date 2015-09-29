package clock

import "fmt"

//TestVersion checks if this code is running on the right test.
const TestVersion = 2

//Clock represents a hour and minute without a date.
//This implementation stores the total minutes of a day
type Clock int

const day = 24 * 60

func convModulo(val int) int {
	return ((val % day) + day) % day
}

//Time returns a clock with the corresponding hour and minutes
func Time(hour, minute int) Clock {
	return Clock(convModulo(day + hour*60 + minute))
}

//Add minutes to this clock
func (c Clock) Add(minutes int) Clock {
	return Clock(convModulo(int(c) + minutes))
}

func (c Clock) String() string {
	h := int(c) / 60
	m := int(c) % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}
