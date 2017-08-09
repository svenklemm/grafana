package metrics

type Gauge interface {
	Metric
	Update(int64)
}

type gauge struct {
	Metric
	name   string
	gauges []Gauge
}

func (g *gauge) Update(v int64) {
	for i, _ := range g.gauges {
		g.gauges[i].Update(v)
	}
}

func (g *gauge) Name() string {
	return g.name
}
