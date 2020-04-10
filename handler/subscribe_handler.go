package handler

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

type SubscribeController struct {}

func(controller SubscribeController) Handle(bot *linebot.Client, event *linebot.Event) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Subscribe")).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}
