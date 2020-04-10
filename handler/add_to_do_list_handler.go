package handler

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

type AddToDoListHandler struct {}

func(handler AddToDoListHandler) Handle(bot *linebot.Client, event *linebot.Event) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Add")).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}