package internal

import (
	"context"

	"github.com/mrbelka12000/optimizer/internal/models"
)

type Adapter interface {
	List(ctx context.Context, pars models.Data) error
}
