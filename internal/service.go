package internal

import "log/slog"

type Service struct {
	next Adapter
	log  *slog.Logger
}

func NewService(adapter Adapter, log *slog.Logger) *Service {
	return &Service{
		next: adapter,
		log:  log,
	}
}
