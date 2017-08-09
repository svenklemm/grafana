package graphite

import (
	"sync/atomic"

	"github.com/grafana/grafana/pkg/metrics"
)

type GraphiteClient struct {
}

func (gc *GraphiteClient) RegCounter(meta *metrics.MetricMeta) metrics.Counter {
	return NewCounter2(meta)
}

func (gc *GraphiteClient) RegGauge(meta *metrics.MetricMeta) metrics.Gauge {
	return NewGauge(meta)
}

func (gc *GraphiteClient) RegTimer(meta *metrics.MetricMeta) metrics.Timer {
	return NewTimer(meta)
}

// Counters hold an int64 value that can be incremented and decremented.
type Counter interface {
	metrics.Metric

	Clear()
	Count() int64
	Dec(int64)
	Inc(int64)
}

// NewCounter constructs a new StandardCounter.
func NewCounter(meta *metrics.MetricMeta) Counter {
	return &StandardCounter{
		MetricMeta: meta,
		count:      0,
	}
}

func NewCounter2(meta *metrics.MetricMeta) Counter {
	cr := &StandardCounter{
		MetricMeta: meta,
		count:      0,
	}
	MetricStats.Register(cr)
	return cr
}

func RegCounter(name string, tagStrings ...string) Counter {
	cr := NewCounter(metrics.NewMetricMeta(name, tagStrings))
	MetricStats.Register(cr)
	return cr
}

// StandardCounter is the standard implementation of a Counter and uses the
// sync/atomic package to manage a single int64 value.
type StandardCounter struct {
	count int64 //Due to a bug in golang the 64bit variable need to come first to be 64bit aligned. https://golang.org/pkg/sync/atomic/#pkg-note-BUG
	*metrics.MetricMeta
}

// Clear sets the counter to zero.
func (c *StandardCounter) Clear() {
	atomic.StoreInt64(&c.count, 0)
}

// Count returns the current count.
func (c *StandardCounter) Count() int64 {
	return atomic.LoadInt64(&c.count)
}

// Dec decrements the counter by the given amount.
func (c *StandardCounter) Dec(i int64) {
	atomic.AddInt64(&c.count, -i)
}

// Inc increments the counter by the given amount.
func (c *StandardCounter) Inc(i int64) {
	atomic.AddInt64(&c.count, i)
}

func (c *StandardCounter) Snapshot() metrics.Metric {
	return &StandardCounter{
		MetricMeta: c.MetricMeta,
		count:      c.count,
	}
}
