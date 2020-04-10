package handler

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

type HelpToDoListHandler struct {}

func (handler HelpToDoListHandler) Handle(bot *linebot.Client, event *linebot.Event) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Help")).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}
