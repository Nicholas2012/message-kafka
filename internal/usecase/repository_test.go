package usecase

import (
	"context"

	"testkafka/internal/models"
)

type repositoryMock struct {
	CreateMessageFn func(ctx context.Context, message *models.Message) error
	UpdateMessageFn func(ctx context.Context, message *models.Message) error
	GetMessageFn    func(ctx context.Context, id int) (*models.Message, error)
	StatisticsFn    func(ctx context.Context) (*models.Statistics, error)
}

func (r *repositoryMock) CreateMessage(ctx context.Context, message *models.Message) error {
	if r.CreateMessageFn == nil {
		return nil
	}
	return r.CreateMessageFn(ctx, message)
}

func (r *repositoryMock) UpdateMessage(ctx context.Context, message *models.Message) error {
	if r.UpdateMessageFn == nil {
		return nil
	}
	return r.UpdateMessageFn(ctx, message)
}

func (r *repositoryMock) GetMessage(ctx context.Context, id int) (*models.Message, error) {
	if r.GetMessageFn == nil {
		return nil, nil
	}
	return r.GetMessageFn(ctx, id)
}

func (r *repositoryMock) Statistics(ctx context.Context) (*models.Statistics, error) {
	if r.StatisticsFn == nil {
		return nil, nil
	}
	return r.StatisticsFn(ctx)
}
