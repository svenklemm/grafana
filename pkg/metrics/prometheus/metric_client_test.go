package prometheus

import (
	"testing"

	"github.com/grafana/grafana/pkg/metrics"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMetricCounters(t *testing.T) {

	Convey("Can create prometheus counter", t, func() {
		mc := metrics.MetricFactories{}
		Init(&PrometheusMetricSettings{}, mc)

		counter := mc.RegCounter("test_counter", "tag1", "value1")
		counter.Inc(1)
	})
}
