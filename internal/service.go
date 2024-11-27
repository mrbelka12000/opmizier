package internal

import (
	"log/slog"
)

// service simple business logic implementer
type service struct {
	next adapter
	log  *slog.Logger
}

func newService(adapter adapter, log *slog.Logger) *service {
	return &service{
		next: adapter,
		log:  log,
	}
}
