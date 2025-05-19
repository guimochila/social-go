// Copyright (c) 2025. guimochila.com. Continuous Learning.

package db

import (
	"context"
	"database/sql"
	"time"
)

// Context Timeout for Database.
const TIMEOUT = 5

func New(dsn string, maxIdleTime time.Duration, maxOpenConns, maxIdleConns int) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxIdleConns)
	db.SetConnMaxIdleTime(maxIdleTime)
	db.SetMaxIdleConns(maxIdleConns)

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
