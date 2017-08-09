package metrics

import "time"

type Timer interface {
	Metric
	Update(time.Duration)
	UpdateSince(time.Time)
}

type timer struct {
	Metric
	name   string
	timers []Timer
}

func (t *timer) Update(duration time.Duration) {
	for i, _ := range t.timers {
		t.timers[i].Update(duration)
	}
}

func (t *timer) UpdateSince(time time.Time) {
	for i, _ := range t.timers {
		t.timers[i].UpdateSince(time)
	}
}

func (t *timer) Name() string {
	return t.name
}
