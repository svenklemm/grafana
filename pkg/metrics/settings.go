package metrics

import ini "gopkg.in/ini.v1"

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

func ParseMetricSettings(settingsFile *ini.File) *MetricSettings {
	var settings = &MetricSettings{
		Enabled:         false,
		IntervalSeconds: 10,
	}

	var section, err = settingsFile.GetSection("metrics")
	if err != nil {
		metricsLogger.Crit("Unable to find metrics config section", "error", err)
		return nil
	}

	settings.Enabled = section.Key("enabled").MustBool(false)
	settings.IntervalSeconds = section.Key("interval_seconds").MustInt64(10)

	return settings
}
