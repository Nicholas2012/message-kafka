package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	goose "github.com/pressly/goose/v3"
	"testkafka/migrations"
)

func New(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func ApplyMigrations(db *sql.DB) error {
	goose.SetBaseFS(migrations.Migrations)
	if err := goose.Up(db, "."); err != nil {
		return err
	}
	return nil
}
