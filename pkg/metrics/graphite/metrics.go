package graphite

import "github.com/grafana/grafana/pkg/metrics"

var UseNilMetrics bool

type MetricPublisher interface {
	Publish(metrics []metrics.Metric)
}
