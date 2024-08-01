package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"testkafka/internal/models"
)

type serviceMock struct {
	createMessageFn func(ctx context.Context, message string) error
	statisticsFn    func(ctx context.Context) (*models.Statistics, error)
}

func (m *serviceMock) CreateMessage(ctx context.Context, message string) error {
	return m.createMessageFn(ctx, message)
}

func (m *serviceMock) Statistics(ctx context.Context) (*models.Statistics, error) {
	return m.statisticsFn(ctx)
}

func setup(t *testing.T) (*httptest.Server, *serviceMock) {
	mux := http.NewServeMux()
	sm := &serviceMock{}

	New(sm).AddRoutes(mux)

	srv := httptest.NewServer(mux)

	return srv, sm
}
