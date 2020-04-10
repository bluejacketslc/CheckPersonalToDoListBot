package controller

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

type EventTypeFollowController struct {}

func (controller *EventTypeFollowController) Execute(bot *linebot.Client, event *linebot.Event) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Hello, New User")).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}