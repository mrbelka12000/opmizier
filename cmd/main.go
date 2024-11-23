package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	"github.com/mrbelka12000/optimizer/internal"
	"github.com/mrbelka12000/optimizer/internal/repository"
	"github.com/mrbelka12000/optimizer/pkg/database"
	"github.com/mrbelka12000/optimizer/pkg/redis"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("service_name", "optimizer")

	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file", err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		log.Error("cannot connect to database", err)
		return
	}
	defer db.Close()

	redisCLI, err := redis.New()
	if err != nil {
		log.Error("cannot connect to redis", err)
		return
	}
	repo := repository.New(db)

	err = internal.Run(redisCLI, repo, log)
	if err != nil {
		log.Error("cannot start app", err)
		return
	}

	log.Info("Shutting down...")
}
