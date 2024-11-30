package timezone

import "time"

func Init(timezone string) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		panic("Invalid timezone: " + timezone)
	}
	time.Local = loc
}
