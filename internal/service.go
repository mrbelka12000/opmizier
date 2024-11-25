package internal

import (
	"log/slog"
)

// Service simple business logic implementer
type Service struct {
	next adapter
	log  *slog.Logger
}

func NewService(adapter adapter, log *slog.Logger) *Service {
	return &Service{
		next: adapter,
		log:  log,
	}
}
