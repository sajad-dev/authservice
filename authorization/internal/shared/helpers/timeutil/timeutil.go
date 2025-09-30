package timeutil

import "time"

func TokenExpire() time.Time {
	return time.Now().AddDate(0, 0, 1)
}
