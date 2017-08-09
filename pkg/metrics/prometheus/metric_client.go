package prometheus

import (
	"strings"
	"time"

	"github.com/grafana/grafana/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusClient struct {
}

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

	gauge prometheus.GaugeVec
}

func (pg *promGauge) Snapshot() metrics.Metric {
	return nil
}

func (pg *promGauge) Update(v int64) {

}

type promTimer struct {
	*metrics.MetricMeta

	timer prometheus.Timer
}

func (pt *promTimer) Update(v time.Duration) {

}

func (pt *promTimer) UpdateSince(v time.Time) {

}

func (pt *promTimer) Snapshot() metrics.Metric {
	return nil
}

func (gc *PrometheusClient) RegCounter(meta *metrics.MetricMeta) metrics.Counter {
	cv := prometheus.NewCounter(prometheus.CounterOpts{
		Name:        prometheusIfyName(meta.Name()) + "_total",
		Help:        meta.Name(),
		ConstLabels: prometheus.Labels(meta.GetTagsCopy()),
	})

	prometheus.MustRegister(cv)

	return &promCounter{
		counter: cv,
	}
}

func (gc *PrometheusClient) RegGauge(meta *metrics.MetricMeta) metrics.Gauge {
	return &promGauge{}
}

func (gc *PrometheusClient) RegTimer(meta *metrics.MetricMeta) metrics.Timer {
	return &promTimer{}
}

func prometheusIfyName(name string) string {
	return strings.Replace(name, ".", "_", -1)
}
