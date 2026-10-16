package telegramapi

import (
	"context"
	"errors"
	"tgFolder/intern/domain"
	"time"

	"gopkg.in/telebot.v3"
	botapi "gopkg.in/telebot.v3"
)

type ServiceInterface interface {
	CreateUser(ctx context.Context, telegramID int64, name string) error
}
type Bot struct {
	Service     ServiceInterface
	TelegramAPI *botapi.Bot
}

func NewBot(service ServiceInterface, token string) (*Bot, error) {
	b, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}
	return &Bot{Service: service, TelegramAPI: b}, nil
}

func (b *Bot) RegisterHandlers() {
	b.TelegramAPI.Handle("/start", b.startChat)
}

func (b *Bot) startChat(c telebot.Context) error {
	ctx := context.Background()
	err := b.Service.CreateUser(ctx, c.Sender().ID, c.Sender().FirstName)
	if err != nil {
		if errors.Is(err, domain.ErrUserExists) {
			return c.Send("Товарищ! Ты уже зарегестрирован!")
		}
		return c.Send("Товарищ! Произошла непредвиденная поломка. Передай генсеку(ссылка в шапке) текст ошибки: " + err.Error())
	}
	return c.Send("Привет, товарищ Комсомолец! Я коммунистическая машина для папок времен СССР.\n" +
		"Команды:\n" +
		"/add <папка> — добавить файл\n" +
		"/get <папка> — получить файлы\n" +
		"/rename <старая> <новая> — переименовать\n" +
		"/delete <папка> — удалить папку")

}
