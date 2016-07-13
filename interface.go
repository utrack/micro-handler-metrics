/*package micrometer provides a HandlerWrapper which reports RPC requests'
timings and counts via provided go-platform/metrics.Metrics spawner.

It reports rpc_response_time_ms counter, rpc_response_time histogram buckets
and rpc_errors counter. Every metric is divided by tags "service" and "method".
*/
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
	return w.wrapWithMetrics
}
