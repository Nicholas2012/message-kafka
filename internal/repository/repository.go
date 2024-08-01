package repository

import (
	"context"
	"database/sql"
	"testkafka/internal/models"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateMessage(ctx context.Context, message *models.Message) error {
	query := `INSERT INTO messages (message) 
		VALUES ($1)
		RETURNING id`

	row := r.db.QueryRowContext(ctx, query, message.Message)
	if err := row.Scan(&message.ID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateMessage(ctx context.Context, message *models.Message) error {
	query := `UPDATE messages SET sent_at = $1 WHERE id = $2`

	_, err := r.db.ExecContext(ctx, query, message.SentAt, message.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Statistics(ctx context.Context) (*models.Statistics, error) {
	query := `SELECT 
    	COUNT(*) AS total,
    	COUNT(sent_at) AS processed
	FROM 
    	messages`

	var stat models.Statistics
	row := r.db.QueryRowContext(ctx, query)
	if err := row.Scan(&stat.Total, &stat.Processed); err != nil {
		return nil, err
	}

	return &stat, nil
}

func (r *Repository) GetMessage(ctx context.Context, id int) (*models.Message, error) {
	query := `SELECT id, message, created_at, sent_at FROM messages WHERE id = $1`

	var m models.Message
	row := r.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&m.ID, &m.Message, &m.CreatedAt, &m.SentAt); err != nil {
		return nil, err
	}

	return &m, nil
}
