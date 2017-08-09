package metrics

type Counter interface {
	Metric
	Inc(int64)
}

type counter struct {
	Metric
	name     string
	counters []Counter
}

func (c *counter) Inc(v int64) {
	for i, _ := range c.counters {
		c.counters[i].Inc(v)
	}
}

func (c *counter) Name() string {
	return c.name
}
