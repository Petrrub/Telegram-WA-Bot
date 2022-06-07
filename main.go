package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
	_ "net/http"
	"strings"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5533440682:AAEkizv_07B9ZY_z1IQDhlahi6lbJo8jPXQ")
	if err != nil {
		log.Panic(err)
	}

	//photoBytes := []byte(response)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if strings.Contains(update.Message.Text, "start") {
				msgfirst := tgbotapi.NewMessage(update.Message.Chat.ID, "Бот запущен")
				bot.Send(msgfirst)
			} else {
				msgfirst := tgbotapi.NewMessage(update.Message.Chat.ID, "Обработка...")
				bot.Send(msgfirst)
				msg := tgbotapi.NewPhoto(update.Message.Chat.ID, GetUrlFile(strings.ReplaceAll(update.Message.Text, "+", "%2B")))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}

func GetUrlFile(message string) tgbotapi.FileBytes {
	url := "http://api.wolframalpha.com/v1/simple?appid=73QR56-W8XYL434X9&i=" + message + "&layout=labelbar"
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	//var b bytes.Buffer
	phBytes, _ := io.ReadAll(response.Body)

	//b.Write([]bytes(response.ContentLength))
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "gif",
		Bytes: phBytes,
	}
	return photoFileBytes
}
