package service

import (
	"context"
	"tgFolder/intern/domain"
	"time"
)

type RepoInterface interface {
	CreateUser(ctx context.Context, u domain.User) error
}

type Service struct {
	repo RepoInterface
}

func NewService(repo RepoInterface) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, telegramID int64, name string) error {
	if len(name) > 100 {
		return domain.ErrTooLongName
	}
	user := domain.User{
		Telegram_id: telegramID,
		Name:        name,
		CreatedAt:   time.Now().UTC(),
	}
	if err := s.repo.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}
