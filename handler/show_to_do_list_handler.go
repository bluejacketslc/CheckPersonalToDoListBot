package handler

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

type ShowToDoListController struct {}

func(controller ShowToDoListController) Handle(bot *linebot.Client, event *linebot.Event) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Show")).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}