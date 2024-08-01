package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"testkafka/internal/models"

	"github.com/stretchr/testify/require"
)

func TestStatistics_OK(t *testing.T) {
	srv, sm := setup(t)

	sm.statisticsFn = func(_ context.Context) (*models.Statistics, error) {
		return &models.Statistics{
			Total:     10,
			Processed: 5,
		}, nil
	}

	req, err := http.NewRequest(http.MethodGet, srv.URL+"/statistics", nil)
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, res.StatusCode)

	var resp StatisticsResponse
	require.NoError(t, json.NewDecoder(res.Body).Decode(&resp))

	require.Equal(t, 10, resp.Total)
	require.Equal(t, 5, resp.Processed)
}

func TestStatistics_InternalServerError(t *testing.T) {
	srv, sm := setup(t)

	sm.statisticsFn = func(_ context.Context) (*models.Statistics, error) {
		return nil, errors.New("some error")
	}

	req, err := http.NewRequest(http.MethodGet, srv.URL+"/statistics", nil)
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	require.Equal(t, http.StatusInternalServerError, res.StatusCode)

	var resp ErrorResponse
	require.NoError(t, json.NewDecoder(res.Body).Decode(&resp))

	require.NotEmpty(t, resp.Error)
	require.Contains(t, resp.Error, "some error")
}
