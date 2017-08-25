package prometheus

import (
	"strings"

	"github.com/grafana/grafana/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusClient struct{}

func (gc *PrometheusClient) RegCounter(meta *metrics.MetricMeta) metrics.Counter {
	cv := prometheus.NewCounter(prometheus.CounterOpts{
		Name:        promifyName(meta.Name()) + "_total",
		Help:        meta.Description(),
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
		Help:        meta.Description(),
		ConstLabels: prometheus.Labels(meta.GetTagsCopy()),
	})

	prometheus.MustRegister(gauge)

	return &promGauge{
		gauge: gauge,
	}
}

func (gc *PrometheusClient) RegTimer(meta *metrics.MetricMeta) metrics.Timer {
	timer := prometheus.NewSummary(prometheus.SummaryOpts{
		Name:        promifyName(meta.Name()) + "_seconds",
		Help:        meta.Description(),
		ConstLabels: prometheus.Labels(meta.GetTagsCopy()),
	})

	prometheus.MustRegister(timer)

	return &promTimer{
		summary: timer,
	}
}

func promifyName(name string) string {
	return strings.Replace(name, ".", "_", -1) //prometheus does not support '.' in the metric name
}
