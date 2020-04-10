package handler

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

type ShowToDoListHandler struct {}

func(handler ShowToDoListHandler) Handle(bot *linebot.Client, event *linebot.Event) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Show")).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}