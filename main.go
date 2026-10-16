package main

import (
	"context"
	"fmt"
	"log"
	"os"
	dbconn "tgFolder/intern/db_conn"
	"tgFolder/intern/repository"
	"tgFolder/intern/service"
	telegramapi "tgFolder/intern/telegram_api"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	ctx := context.Background()
	key := os.Getenv("TG_TOKEN")
	connStr := os.Getenv("CONN_STR")
	m, err := migrate.New("file://migrations", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	pool, err := dbconn.NewConnection(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	repo := repository.NewRepository(pool)
	serv := service.NewService(repo)
	bot, err := telegramapi.NewBot(serv, key)
	if err != nil {
		log.Fatal(err)
	}
	bot.RegisterHandlers()
	bot.TelegramAPI.Start()
	fmt.Println("bot successfully started")
}
