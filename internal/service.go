package internal

import (
	"log/slog"

	"github.com/gorilla/schema"
)

// Service simple business logic implementer
type Service struct {
	next    adapter
	decoder *schema.Decoder
	log     *slog.Logger
}

func NewService(adapter adapter, log *slog.Logger) *Service {
	return &Service{
		next:    adapter,
		decoder: schema.NewDecoder(),
		log:     log,
	}
}
