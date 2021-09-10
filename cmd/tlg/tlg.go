package tlg

import (
	//"fmt"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zmb3/spotify/v2"
)

var (
	TG_TOKEN = "1990841068:AAE8g-yvlRb1yfQBP_sd5ECd4-ojuy3VI2M"
	DEBUG    = true
)

func Init(cl *spotify.Client, usr *spotify.PrivateUser) {
	bot, err := tgbotapi.NewBotAPI(TG_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			// ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg, err := process(bot, &update, cl, usr)

		if err != nil {
			str := fmt.Sprintf("О ужос, случилась ошибка: %s", err.Error())
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, str)
			bot.Send(msg)
		} else {
			bot.Send(msg)
		}

	}

	// return bot, err

}
