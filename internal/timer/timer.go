package timer

import (
	"time"
)

type Timer struct {
	Duration     time.Duration
	StartTime    time.Time
	IsRunning    bool
	TotalElapsed time.Duration
}

func NewTimer(duration time.Duration) *Timer {
	return &Timer{
		Duration: duration,
	}
}

func (t *Timer) Start() {
	if !t.IsRunning {
		t.StartTime = time.Now()
		t.IsRunning = true
	}
}

func (t *Timer) Stop() {
	if t.IsRunning {
		t.TotalElapsed += time.Since(t.StartTime)
		t.IsRunning = false
	}
}

func (t *Timer) Reset() {
	t.TotalElapsed = 0
	t.IsRunning = false
}

func (t *Timer) Elapsed() time.Duration {
	if t.IsRunning {
		return t.TotalElapsed + time.Since(t.StartTime)
	}
	return t.TotalElapsed
}

func (t *Timer) Remaining() time.Duration {
	return t.Duration - t.Elapsed()
}
