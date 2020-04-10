package main

import (
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
	"todoreminder/constants"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func initBot() *linebot.Client {
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err.Error())
	}

	return bot
}

func setListeners(bot *linebot.Client, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, event := range events {
		constants.EventListeners[event.Type].Execute(bot, event)
	}
}

func main() {
	loadEnv()
	bot := initBot()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		setListeners(bot, r)
	})

	err := http.ListenAndServe(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}