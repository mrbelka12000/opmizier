package internal

import (
	"context"
	"log/slog"
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
	return &metricsMiddleware{
		next: next,
		log:  log,
	}
}

func (m *metricsMiddleware) List(ctx context.Context, query string, args []any) error {
	return m.next.List(ctx, query, args)
}
