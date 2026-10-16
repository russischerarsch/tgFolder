package domain

import "time"

type User struct {
	ID          int64
	Telegram_id int64
	Name        string
	CreatedAt   time.Time
}
