package tlg

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zmb3/spotify/v2"
)

var (
	mainMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("📟"),
			tgbotapi.NewKeyboardButton("➕"),
			// tgbotapi.NewKeyboardButton("showall"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("⏮️"),
			tgbotapi.NewKeyboardButton("⏯️"),
			tgbotapi.NewKeyboardButton("⏭️"),
		),
		// tgbotapi.NewKeyboardButtonRow(
		// 	tgbotapi.NewKeyboardButton("plus"),
		// 	tgbotapi.NewKeyboardButton("help"),
		// 	tgbotapi.NewKeyboardButton("minus"),
		// ),
	)
	waitingForSong = false
)

func process(bot *tgbotapi.BotAPI, update *tgbotapi.Update, cl *spotify.Client, usr *spotify.PrivateUser) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Открой меню написав мне слово МАША, и тыкай кнопки на супер мега клавиатуре")
	if waitingForSong {
		str := update.Message.Text
		results, err := cl.Search(context.Background(), str, spotify.SearchTypeTrack)
		if err != nil {
			return msg, err
		}

		// handle album results
		if results.Tracks != nil && results.Tracks.Tracks != nil && (len(results.Tracks.Tracks) > 0) {
			item := results.Tracks.Tracks[0]
			msg.Text = fmt.Sprintf("%s - %s", item.Name, item.Artists[0].Name)
			msg.ReplyMarkup = mainMenuKeyboard
			cl.QueueSong(context.Background(), item.ID)
		} else {
			msg.Text = fmt.Sprintf("Ну ты походу пьяная свинья, ничего не найдено по твоему запросу")
		}

	}
	waitingForSong = false

	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			msg.ReplyMarkup = mainMenuKeyboard
			msg.Text = "Добро пожаловать на музыкальную вечеринку!\nИспользуй кнопки снизу для управляния\n 📟 - посмотреть че играет сейчас\n ➕ - добавить трек в очередь\n ну а снизу и так понятно. Удачи в алкомузлотрипе:) "
		case "help":
			msg.Text = "Command use\n\t /help - to see this again\n\t /start - to get super-mega-menu\n Используй кнопки снизу для управляния\n 📟 - посмотреть че играет сейчас\n ➕ - добавить трек в очередь\n ну а снизу и так понятно. Удачи в алкомузлотрипе:) "
		default:
			msg.Text = "I don't know that command use  /help /start"
		}
	}
	switch update.Message.Text {
	case "📟":
		track, err := cl.PlayerCurrentlyPlaying(context.Background())
		if err != nil {
			return msg, err
		}
		fmt.Printf("\n%+v\n", track)
		msg.Text = fmt.Sprintf("%s - %s", track.Item.Artists[0].Name, track.Item.Name)
	case "showall":
		//todo
	case "start":
		msg.ReplyMarkup = mainMenuKeyboard
	case "МАША":
		msg.ReplyMarkup = mainMenuKeyboard
		msg.Text = "Красссавчик!(или красавица)"
	case "⏭️":
		cl.Next(context.Background())
		msg.Text = "Включаю следующую, но знай ктото грустит изза того что пропустил его песню"
	case "⏮️":
		cl.Previous(context.Background())
		msg.Text = "Включаю предыдущую, но знай ктото радуется изза того что повторил его песню (если ты посторянно повторяешь свою песню, знай она всем надоела)"
	case "⏯️":
		msg.Text = "Пауза или плай, хозяин поленился писать вариант под оба случая, поэтому так 8=э"
		state, err := cl.PlayerState(context.Background())
		if err != nil {
			return msg, err
		}
		if state.Playing {
			cl.Pause(context.Background())
		} else {
			cl.Play(context.Background())
		}
	case "➕":
		msg.Text = "Send song in format example \"Паша техник - нужен ксанокс\""
		waitingForSong = true
		// case "plus":
		// 	state, err := cl.PlayerState(context.Background())
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	volume := state.Device.Volume
		// 	if (volume + 10) >= 100 {
		// 		volume = 100
		// 		msg.Text = "Volume set to 100%"
		// 	} else {
		// 		volume += 10
		// 		msg.Text = fmt.Sprintf("Volume set to %v%%", volume)
		// 	}
		// 	err=cl.Volume(context.Background(), volume)
		// 	if err != nil {
		// 	  panic(err)
		// 	}
		// case "help":
		// 	msg.Text = "help message"
		// case "minus":
		// 	state, err := cl.PlayerState(context.Background())
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	volume := state.Device.Volume
		// 	if (volume - 10) <= 0 {
		// 		volume = 0
		// 		msg.Text = "Volume set to 0%"
		// 	} else {
		// 		volume -= 10
		// 		msg.Text = fmt.Sprintf("Volume set to %v%%", volume)
		// 	}
		// 	err=cl.Volume(context.Background(), volume)
		// 	if err != nil {
		// 	  panic(err)
		// 	}
	}

	return msg, nil
}
