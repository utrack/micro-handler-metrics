package micrometer

import (
	"github.com/micro/go-micro/server"
	"github.com/micro/go-platform/metrics"
)

// NewWrapper returns the HandlerWrapper which instruments the handlers
// using given metrics reporter.
func NewWrapper(mSpawner metrics.Metrics) server.HandlerWrapper {
	w := &w{
		timeCount: mSpawner.Counter("rpc_response_time_ms"),
		timeHisto: mSpawner.Histogram("rpc_response_time"),
		errCount:  mSpawner.Counter("rpc_errors"),
	}
	return w.WrapWithMetrics
}
