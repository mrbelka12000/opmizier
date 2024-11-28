package internal

import (
	"context"
)

type adapter interface {
	List(ctx context.Context, query string) error
}
