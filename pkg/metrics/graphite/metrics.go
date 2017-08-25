package graphite

import "github.com/grafana/grafana/pkg/metrics"

var MetricStats Registry
var UseNilMetrics bool

type MetricSettings struct {
	Enabled         bool
	IntervalSeconds int64

	Publishers []MetricPublisher
}

type MetricPublisher interface {
	Publish(metrics []metrics.Metric)
}
