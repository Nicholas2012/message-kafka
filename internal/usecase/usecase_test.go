package usecase

import (
	"context"
	"testing"
	"testkafka/internal/broker"
	"testkafka/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateMessage_OK(t *testing.T) {
	s, repo, b := setup(t)

	repo.CreateMessageFn = func(ctx context.Context, message *models.Message) error {
		assert.Equal(t, "1234 567890", message.Message)

		message.ID = 9899

		return nil
	}

	b.ProduceWithDeliveryFn = func(data []byte, fn broker.OnDeliveryFunc) error {
		assert.Equal(t, `{"ID":9899,"Message":"1234 567890"}`, string(data))

		fn(nil)

		return nil
	}

	repo.UpdateMessageFn = func(ctx context.Context, message *models.Message) error {
		assert.Equal(t, 9899, message.ID)
		assert.Equal(t, "1234 567890", message.Message)

		return nil
	}

	err := s.CreateMessage(context.TODO(), "1234 567890")
	require.NoError(t, err)
}

func setup(_ *testing.T) (*Service, *repositoryMock, *brokerMock) {
	repo := &repositoryMock{}
	bm := &brokerMock{}
	return New(repo, bm), repo, bm
}

type brokerMock struct {
	ProduceWithDeliveryFn func(data []byte, fn broker.OnDeliveryFunc) error
}

func (b *brokerMock) ProduceWithDelivery(data []byte, fn broker.OnDeliveryFunc) error {
	if b.ProduceWithDeliveryFn == nil {
		return nil
	}

	return b.ProduceWithDeliveryFn(data, fn)
}
