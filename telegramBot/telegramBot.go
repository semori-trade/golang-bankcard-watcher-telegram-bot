package telegramBot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/semori-trade/golang-bankcard-watcher-telegram-bot/telegramBot/internal/keyboard"
)

type Options struct {
	Debug bool
}

func NewTelegramBot(token string, options Options) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("1")

	fmt.Println(options)

	if options.Debug {
		bot.Debug = true

		log.Printf("Authorized on account %s", bot.Self.UserName)

		listenUpdates(bot)
	}

	return bot, nil
}

func listenUpdates(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			msg.ReplyMarkup = keyboard.NumericKeyboard

			bot.Send(msg)
		}
	}
}
