package internal

import (
	"context"

	"github.com/mrbelka12000/optimizer/internal/models"
)

type adapter interface {
	List(ctx context.Context, pars models.Request) error
}
