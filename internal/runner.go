package internal

import (
	"log/slog"
	"net/http"

	"github.com/go-redis/redis"

	"github.com/mrbelka12000/optimizer/internal/repository"
	"github.com/mrbelka12000/optimizer/pkg/server"
)

func Run(
	client *redis.Client,
	repo *repository.Repo,
	log *slog.Logger,
) error {

	cache := newCache(client, repo, log)
	metrics := newMetricsMiddleware(cache, log)
	srv := newService(metrics, log)

	mux := http.NewServeMux()
	srv.RegisterHandlers(mux)

	httpServer := server.New(mux)

	httpServer.Start()

	log.Info("Server started")

	return <-httpServer.Ch()
}
