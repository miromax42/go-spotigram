package tlg

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	TG_TOKEN = "1990841068:AAG_lMW107e-21sAIaTTW2A97YsYRwZNJek"
	DEBUG    = true
)

func Init() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(TG_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = DEBUG

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := process(&update)
		fmt.Printf("-------endend----------%v\n", msg)

		bot.Send(msg)

	}

	return bot, err

}
