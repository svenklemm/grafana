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
	pt.summary.Observe(float64(v.Seconds()))
}

func (pt *promTimer) UpdateSince(v time.Time) {
	pt.summary.Observe(float64(time.Since(v).Seconds()))
}

func (pt *promTimer) Snapshot() metrics.Metric {
	return nil
}

func (gc *PrometheusClient) RegCounter(meta *metrics.MetricMeta) metrics.Counter {
	cv := prometheus.NewCounter(prometheus.CounterOpts{
		Name:        promifyName(meta.Name()) + "_total",
		Help:        meta.Name(),
		ConstLabels: prometheus.Labels(meta.GetTagsCopy()),
	})

	prometheus.MustRegister(cv)

	return &promCounter{
		counter: cv,
	}
}

func (gc *PrometheusClient) RegGauge(meta *metrics.MetricMeta) metrics.Gauge {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        promifyName(meta.Name()),
		Help:        meta.Name(),
		ConstLabels: prometheus.Labels(meta.GetTagsCopy()),
	})

	prometheus.MustRegister(gauge)

	return &promGauge{
		gauge: gauge,
	}
}

func (gc *PrometheusClient) RegTimer(meta *metrics.MetricMeta) metrics.Timer {
	timer := prometheus.NewSummary(prometheus.SummaryOpts{
		Name:        promifyName(meta.Name()),
		Help:        meta.Name(),
		ConstLabels: prometheus.Labels(meta.GetTagsCopy()),
	})

	prometheus.MustRegister(timer)

	return &promTimer{
		summary: timer,
	}
}

func promifyName(name string) string {
	return strings.Replace(name, ".", "_", -1)
}
