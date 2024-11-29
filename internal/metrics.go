package internal

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	lblService = "service"
	lblMethod  = "method"
)

type metricsMiddleware struct {
	next    adapter
	latency *stdprometheus.GaugeVec
	log     *slog.Logger
}

func newMetricsMiddleware(next adapter, log *slog.Logger) *metricsMiddleware {
	stdprometheus.MustRegister(stdServerSLALatency)

	return &metricsMiddleware{
		next:    next,
		latency: stdServerSLALatency,
		log:     log,
	}
}

var (
	stdServerSLALatency = stdprometheus.NewGaugeVec(
		stdprometheus.GaugeOpts{
			Namespace: "server",
			Name:      "http_request_duration_milly_seconds",
		}, []string{lblService, lblMethod},
	)
)

func (m *metricsMiddleware) List(ctx context.Context, query string) error {
	start := time.Now()
	id := uuid.New().String()

	m.log.Info(fmt.Sprintf("Enter: %s", id))
	defer func() {
		m.write("List", time.Since(start).Seconds())
		m.log.Info(fmt.Sprintf("Finished: %s , spent: %.5f seconds", id, time.Since(start).Seconds()))
	}()

	return m.next.List(ctx, query)
}

func (m *metricsMiddleware) write(method string, duration float64) {
	m.latency.WithLabelValues("optimizer", method).Set(duration)
}
