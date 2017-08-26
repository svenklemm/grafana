package metrics

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/plugins"
	"github.com/grafana/grafana/pkg/setting"
)

func sendUsageStats() {
	if !setting.ReportingEnabled {
		return
	}

	metricsLogger.Debug("Sending anonymous usage stats to stats.grafana.org")

	version := strings.Replace(setting.BuildVersion, ".", "_", -1)

	metrics := map[string]interface{}{}
	report := map[string]interface{}{
		"version": version,
		"metrics": metrics,
	}

	statsQuery := models.GetSystemStatsQuery{}
	if err := bus.Dispatch(&statsQuery); err != nil {
		metricsLogger.Error("Failed to get system stats", "error", err)
		return
	}

	metrics["stats.dashboards.count"] = statsQuery.Result.Dashboards
	metrics["stats.users.count"] = statsQuery.Result.Users
	metrics["stats.orgs.count"] = statsQuery.Result.Orgs
	metrics["stats.playlist.count"] = statsQuery.Result.Playlists
	metrics["stats.plugins.apps.count"] = len(plugins.Apps)
	metrics["stats.plugins.panels.count"] = len(plugins.Panels)
	metrics["stats.plugins.datasources.count"] = len(plugins.DataSources)
	metrics["stats.alerts.count"] = statsQuery.Result.Alerts

	dsStats := models.GetDataSourceStatsQuery{}
	if err := bus.Dispatch(&dsStats); err != nil {
		metricsLogger.Error("Failed to get datasource stats", "error", err)
		return
	}

	// send counters for each data source
	// but ignore any custom data sources
	// as sending that name could be sensitive information
	dsOtherCount := 0
	for _, dsStat := range dsStats.Result {
		if models.IsKnownDataSourcePlugin(dsStat.Type) {
			metrics["stats.ds."+dsStat.Type+".count"] = dsStat.Count
		} else {
			dsOtherCount += dsStat.Count
		}
	}
	metrics["stats.ds.other.count"] = dsOtherCount

	out, _ := json.MarshalIndent(report, "", " ")
	data := bytes.NewBuffer(out)

	client := http.Client{Timeout: time.Duration(5 * time.Second)}
	go client.Post("https://stats.grafana.org/grafana-usage-report", "application/json", data)
}
