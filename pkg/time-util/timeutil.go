package timeutil

import "time"

var TimeNowUte = func() time.Time {
	return time.Now().UTC()
}
