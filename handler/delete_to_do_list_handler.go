package handler

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

type DeleteToDoListController struct {}

func(controller DeleteToDoListController) Handle(bot *linebot.Client, event *linebot.Event) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Delete")).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}