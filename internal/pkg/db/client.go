package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Client structure
type Client interface {
	Close() error
	DB() *DB
}

type client struct {
	db *DB
}

// NewClient starts client
func NewClient(ctx context.Context, config *pgxpool.Config) (Client, error) {
	dbc, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &client{
		db: &DB{pool: dbc},
	}, nil
}

func (c *client) Close() error {
	if c != nil {
		c.db.Close()
	}

	return nil
}

func (c *client) DB() *DB {
	return c.db
}
