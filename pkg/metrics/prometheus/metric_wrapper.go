package prometheus

import (
	"time"

	"github.com/grafana/grafana/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type promCounter struct {
	*metrics.MetricMeta

	counter prometheus.Counter
}

func (pc *promCounter) Inc(v int64) {
	pc.counter.Add(float64(v))
}

func (pc *promCounter) Snapshot() metrics.Metric {
	return nil
}

type promGauge struct {
	*metrics.MetricMeta

	gauge prometheus.Gauge
}

func (pg *promGauge) Snapshot() metrics.Metric {
	return nil
}

func (pg *promGauge) Update(v int64) {
	pg.gauge.Set(float64(v))
}

type promTimer struct {
	*metrics.MetricMeta

	summary prometheus.Summary
}

func (pt *promTimer) Update(v time.Duration) {
	pt.summary.Observe(v.Seconds())
}

func (pt *promTimer) UpdateSince(v time.Time) {
	pt.summary.Observe(time.Since(v).Seconds())
}

func (pt *promTimer) Snapshot() metrics.Metric {
	return nil
}
