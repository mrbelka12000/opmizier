package internal

import (
	"context"
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	lblService = "service"
	lblMethod  = "method"
)

type metricsMiddleware struct {
	next adapter
	log  *slog.Logger
}

func newMetricsMiddleware(next adapter, log *slog.Logger) *metricsMiddleware {
	prometheus.MustRegister()
	return &metricsMiddleware{
		next: next,
		log:  log,
	}
}

func (m *metricsMiddleware) List(ctx context.Context, query string) error {
	return m.next.List(ctx, query)
}
