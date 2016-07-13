package micrometer

import (
	"github.com/micro/go-micro/server"
	"github.com/micro/go-platform/metrics"
	"golang.org/x/net/context"
	"time"
)

type w struct {
	timeCount metrics.Counter
	timeHisto metrics.Histogram
	errCount  metrics.Counter
}

var _ server.HandlerWrapper = (&w{}).WrapWithMetrics

func (w *w) WrapWithMetrics(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fields := metrics.Fields{"service": req.Service(), "method": req.Method()}

		start := time.Now()
		err := fn(ctx, req, rsp)
		if err != nil {
			w.errCount.WithFields(fields).Incr(1)
		}

		millis := uint64(time.Since(start).Seconds() * 1000)

		w.timeCount.WithFields(fields).Incr(millis)
		w.timeHisto.WithFields(fields).Record(int64(millis))

		return err
	}
}
