package metrics

type MetricFactory interface {
	RegCounter(*MetricMeta) Counter
	RegTimer(*MetricMeta) Timer
	RegGauge(*MetricMeta) Gauge
}

type MetricFactories map[string]MetricFactory

func (mc MetricFactories) RegCounter(name string, tagStrings ...string) Counter {
	mm := NewMetricMeta(name, tagStrings)
	counters := []Counter{}

	for _, v := range mc {
		counters = append(counters, v.RegCounter(mm))
	}

	return &counter{
		counters: counters,
	}
}

func (mc MetricFactories) RegTimer(name string, tagStrings ...string) Timer {
	mm := NewMetricMeta(name, tagStrings)
	timers := []Timer{}

	for _, v := range mc {
		timers = append(timers, v.RegTimer(mm))
	}

	return &timer{
		timers: timers,
	}
}

func (mc MetricFactories) RegGauge(name string, tagStrings ...string) Gauge {
	mm := NewMetricMeta(name, tagStrings)
	gauges := []Gauge{}

	for _, v := range mc {
		gauges = append(gauges, v.RegGauge(mm))
	}

	return &gauge{
		gauges: gauges,
	}
}

type MetricSettings struct {
	Enabled         bool
	IntervalSeconds int64
}
