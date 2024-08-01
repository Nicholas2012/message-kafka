package main

import (
	"log/slog"
	"net/http"
	"os"
	"testkafka/internal/api"
	"testkafka/internal/broker"
	"testkafka/internal/config"
	"testkafka/internal/repository"
	"testkafka/internal/usecase"
	"testkafka/pkg/database"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	slog.Info("Starting server...")

	config := config.New()

	db, err := database.New(config.DatabaseDSN)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	if err := database.ApplyMigrations(db); err != nil {
		slog.Error("setup db connection", "err", err)
		os.Exit(1)
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.KafkaBroker})
	if err != nil {
		slog.Error("Failed to create Kafka producer", "error", err)
		os.Exit(1)
	}
	defer producer.Close()

	var (
		kb   = broker.New(producer, config.KafkaTopic)
		repo = repository.New(db)
		svc  = usecase.New(repo, kb)
		api  = api.New(svc)
	)

	api.AddRoutes(http.DefaultServeMux)

	slog.Info("Server started", "Listen", config.Listen)
	if err := http.ListenAndServe(config.Listen, nil); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
