package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"testkafka/internal/broker"
	"testkafka/internal/models"
)

type Producer interface {
	ProduceWithDelivery(data []byte, onDelivery broker.OnDeliveryFunc) error
}

type Service struct {
	repo     Repository
	producer Producer
}

func New(repo Repository, producer Producer) *Service {
	return &Service{
		repo:     repo,
		producer: producer,
	}
}

func (s *Service) CreateMessage(ctx context.Context, message string) error {
	newMessage := &models.Message{
		Message: message,
	}

	if err := s.repo.CreateMessage(ctx, newMessage); err != nil {
		return fmt.Errorf("create message: %w", err)
	}

	messageBytes, err := json.Marshal(broker.NewMessageEvent(*newMessage))
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	if err := s.producer.ProduceWithDelivery(messageBytes, s.onDelivery(newMessage)); err != nil {
		return fmt.Errorf("produce message: %w", err)
	}

	return nil
}

func (s *Service) Statistics(ctx context.Context) (*models.Statistics, error) {
	stat, err := s.repo.Statistics(ctx)

	if err != nil {
		return nil, fmt.Errorf("statistics: %w", err)
	}

	return stat, nil
}

func (s *Service) onDelivery(msg *models.Message) broker.OnDeliveryFunc {
	return func(err error) {
		if err != nil {
			slog.Error("message delivery failed", "message_id", msg.ID, "error", err)
			return
		}

		msg.SetSentNow()

		if err := s.repo.UpdateMessage(context.Background(), msg); err != nil {
			slog.Error("update message failed", "message_id", msg.ID, "error", err)
		}
	}
}
