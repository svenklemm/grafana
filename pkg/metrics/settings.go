package metrics

import "github.com/grafana/grafana/pkg/log"

type MetricPublisher interface {
	Publish(metrics []Metric)
}

type MetricClient interface {
	RegCounter(*MetricMeta) Counter
	RegTimer(*MetricMeta) Timer
	RegGauge(*MetricMeta) Gauge
}

type MetricClients map[string]MetricClient

func (mc MetricClients) RegCounter(name string, tagStrings ...string) Counter {
	mm := NewMetricMeta(name, tagStrings)
	counters := []Counter{}

	for _, v := range mc {
		counters = append(counters, v.RegCounter(mm))
	}

	return &counter{
		counters: counters,
	}
}

func (mc MetricClients) RegTimer(name string, tagStrings ...string) Timer {
	mm := NewMetricMeta(name, tagStrings)
	timers := []Timer{}

	for _, v := range mc {
		timers = append(timers, v.RegTimer(mm))
	}

	return &timer{
		timers: timers,
	}
}

func (mc MetricClients) RegGauge(name string, tagStrings ...string) Gauge {
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

	Publishers []MetricPublisher
}

var metricsLogger log.Logger = log.New("metrics")
