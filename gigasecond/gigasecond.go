package gigasecond

import "time"

//TestVersion checks if this code is running on the right test.
const TestVersion = 2

//Birthday is my birthday date
var Birthday, _ = time.Parse("2006-01-02", "1977-01-04")

//AddGigasecond returns the informed time plus 1 billion seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1e9) * time.Second)
}
