package graphite

import (
	"time"

	"github.com/grafana/grafana/pkg/log"
)

var metricsLogger log.Logger = log.New("metrics")
var metricPublishCounter int64 = 0

func instrumentationLoop(settings *GraphiteSettings) {
	if !settings.Enabled || len(settings.Publishers) == 0 {
		return
	}

	secondTicker := time.NewTicker(time.Second * time.Duration(settings.IntervalSeconds))

	for {
		select {
		case <-secondTicker.C:
			metrics := MetricStats.GetSnapshots()
			for _, publisher := range settings.Publishers {
				publisher.Publish(metrics)
			}
		}
	}
}
