package graphite

import "github.com/grafana/grafana/pkg/metrics"

type GraphiteMetricFactory struct {
}

func (gc *GraphiteMetricFactory) RegCounter(meta *metrics.MetricMeta) metrics.Counter {
	return NewCounter2(meta)
}

func (gc *GraphiteMetricFactory) RegGauge(meta *metrics.MetricMeta) metrics.Gauge {
	return NewGauge(meta)
}

func (gc *GraphiteMetricFactory) RegTimer(meta *metrics.MetricMeta) metrics.Timer {
	return NewTimer(meta)
}
