package prometheus

import "github.com/grafana/grafana/pkg/metrics"

func Init(clients metrics.MetricFactories) {
	clients["prometheus"] = &PrometheusClient{}
}
