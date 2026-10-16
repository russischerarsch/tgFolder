package dbconn

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection(ctx context.Context, connStr string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		time.Sleep(1 * time.Second)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}
