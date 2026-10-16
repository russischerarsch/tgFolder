package repository

import (
	"context"
	"fmt"
	"tgFolder/intern/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repoDB struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repoDB {
	return &repoDB{db: db}
}

func (r *repoDB) CreateUser(ctx context.Context, u domain.User) error {
	query := `
	INSERT INTO users(telegram_id, name, created_at)
	VALUES($1, $2, $3)
	`
	if _, err := r.db.Exec(ctx, query, u.Telegram_id, u.Name, u.CreatedAt); err != nil {
		return fmt.Errorf("failed to send query, %w", err)
	}
	return nil
}
