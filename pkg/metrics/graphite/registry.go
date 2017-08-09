package graphite

import (
	"sync"

	"github.com/grafana/grafana/pkg/metrics"
)

type Registry interface {
	GetSnapshots() []metrics.Metric
	Register(metric metrics.Metric)
}

// The standard implementation of a Registry is a mutex-protected map
// of names to metrics.
type StandardRegistry struct {
	metrics []metrics.Metric
	mutex   sync.Mutex
}

// Create a new registry.
func NewRegistry() Registry {
	return &StandardRegistry{
		metrics: make([]metrics.Metric, 0),
	}
}

func (r *StandardRegistry) Register(metric metrics.Metric) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.metrics = append(r.metrics, metric)
}

// Call the given function for each registered metric.
func (r *StandardRegistry) GetSnapshots() []metrics.Metric {
	metrics := make([]metrics.Metric, len(r.metrics))
	for i, metric := range r.metrics {
		metrics[i] = metric.Snapshot()
	}
	return metrics
}
