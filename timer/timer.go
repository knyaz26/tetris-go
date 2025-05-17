package timer

import "time"

type Timer struct {
	startTime time.Time
	duration  time.Duration
}

func StartTimer(durationSeconds int) *Timer {
	return &Timer{
		startTime: time.Now(),
		duration:  time.Duration(durationSeconds) * time.Second,
	}
}

func IsTimerFinished(t *Timer) bool {
	return time.Since(t.startTime) >= t.duration
}
