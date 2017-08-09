package prometheus

import (
	"github.com/grafana/grafana/pkg/metrics"
	ini "gopkg.in/ini.v1"
)

type PrometheusMetricSettings struct {
	Enabled bool
}

func NewPrometheusMetricSettings(settingsFile *ini.File) *PrometheusMetricSettings {
	return &PrometheusMetricSettings{Enabled: true}
}

func Init(settings *PrometheusMetricSettings, clients metrics.MetricClients) {
	if settings.Enabled {
		clients["prometheus"] = &PrometheusClient{}
	}
}
