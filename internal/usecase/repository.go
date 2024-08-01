package usecase

import (
	"context"

	"testkafka/internal/models"
)

type Repository interface {
	CreateMessage(ctx context.Context, message *models.Message) error
	UpdateMessage(ctx context.Context, message *models.Message) error
	GetMessage(ctx context.Context, id int) (*models.Message, error)
	Statistics(ctx context.Context) (*models.Statistics, error)
}
