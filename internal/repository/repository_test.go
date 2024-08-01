package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/ory/dockertest/v3"

	"testkafka/internal/models"
	"testkafka/pkg/database"

	"github.com/stretchr/testify/require"
)

func TestRepository(t *testing.T) {
	repo := setup(t)

	message := &models.Message{
		Message: "test",
	}

	t.Run("CreateMessage", func(t *testing.T) {
		err := repo.CreateMessage(context.Background(), message)
		require.NoError(t, err)
		require.NotZero(t, message.ID)

		t.Run("GetMessage", func(t *testing.T) {
			u, err := repo.GetMessage(context.Background(), message.ID)
			require.NoError(t, err)
			require.Equal(t, message.ID, u.ID)
			require.Equal(t, message.Message, u.Message)
			require.NotEmpty(t, u.CreatedAt)
			require.Nil(t, u.SentAt)
		})
	})
}

func TestMigrations(t *testing.T) {
	setup(t)
}

func setup(t *testing.T) *Repository {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err)

	if err := pool.Client.Ping(); err != nil {
		t.Skipf("test will be skipped because docker is not running, ping error: %s", err)
	}

	resource, err := pool.Run("postgres", "latest",
		[]string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_DB=testdb",
		})
	require.NoError(t, err)

	var db *sql.DB

	err = pool.Retry(func() error {
		newDB, err := database.New(fmt.Sprintf("postgres://postgres:secret@localhost:%s/testdb?sslmode=disable", resource.GetPort("5432/tcp")))
		if err != nil {
			return err
		}
		db = newDB
		return nil
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource))
	})

	require.NoError(t, database.ApplyMigrations(db))

	return New(db)
}
