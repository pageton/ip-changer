package scheduler

import "time"

func ScheduleIPChange(changeIPFunc func(), interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		changeIPFunc()
	}
}
