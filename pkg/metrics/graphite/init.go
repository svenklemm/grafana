package graphite

import (
	"github.com/grafana/grafana/pkg/metrics"
	ini "gopkg.in/ini.v1"
)

type GraphiteSettings struct {
	Enabled         bool
	Address         string
	IntervalSeconds int64

	Publishers []MetricPublisher
}

func NewGraphiteSettings(settingsFile *ini.File) *GraphiteSettings {
	var settings = &GraphiteSettings{
		Enabled:    false,
		Publishers: make([]MetricPublisher, 0),
	}

	var section, err = settingsFile.GetSection("metrics")
	if err != nil {
		metricsLogger.Crit("Unable to find metrics config section", "error", err)
		return nil
	}

	settings.Enabled = section.Key("enabled").MustBool(false)
	settings.IntervalSeconds = section.Key("interval_seconds").MustInt64(10)

	if !settings.Enabled {
		return settings
	}

	if graphitePublisher, err := CreateGraphitePublisher(); err != nil {
		metricsLogger.Error("Failed to init Graphite metric publisher", "error", err)
	} else if graphitePublisher != nil {
		metricsLogger.Info("Metrics publisher initialized", "type", "graphite")
		settings.Publishers = append(settings.Publishers, graphitePublisher)
	}

	return settings
}

func Init(settings *GraphiteSettings, metricClients metrics.MetricClients) {
	metricsLogger.Info("setting up metrics using graphite")

	//settings := readSettings(settingsFile)
	initMetricVars(settings)
	//metrics.RegisterMetricClient("graphite", &GraphiteClient{})
	metricClients["graphite"] = &GraphiteClient{}
	go instrumentationLoop(settings)
}
