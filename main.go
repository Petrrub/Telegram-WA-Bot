package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
	_ "net/http"
	"os"
	"strings"
)

func main() {
	//Get tokens
	dat, err := os.ReadFile("token.txt")
	if err != nil {
		log.Panic(err)
	}

	tokens := strings.Split(string(dat), "\n")
	if len(tokens) != 2 || len(tokens[0]) == 0 || len(tokens[1]) == 0 {
		log.Fatal("\nYou must use two tokens without null value. Example:\nTGtoken\nWAToken ")
	}
	tokens[1] = strings.ReplaceAll(tokens[1], "\r", "")

	//Start bot
	bot, err := tgbotapi.NewBotAPI(strings.ReplaceAll(tokens[0], "\r", ""))
	if err != nil {
		log.Panic(err)
	}

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
				msg := tgbotapi.NewPhoto(update.Message.Chat.ID, GetUrlFile(strings.ReplaceAll(update.Message.Text, "+", "%2B"), tokens[1]))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}

func GetUrlFile(message string, tokenWA string) tgbotapi.FileBytes {
	log.Print(tokenWA + tokenWA)
	url := "http://api.wolframalpha.com/v1/simple?appid=" + tokenWA + "&i=" + message + "&layout=labelbar"
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	phBytes, _ := io.ReadAll(response.Body)
	response.Body.Close()
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "gif",
		Bytes: phBytes,
	}
	return photoFileBytes
}
