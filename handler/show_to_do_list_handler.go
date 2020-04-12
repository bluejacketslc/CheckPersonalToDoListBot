package handler

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

var (
	listHelpInstructions =
		"Command Usage:\n" +
		"/list -> Show 5 near deadline To Do Lists\n" +
		"/list [date in yyyy-mm-dd] -> Show all To Do Lists on that date"
)

type ShowToDoListHandler struct {}

func(handler ShowToDoListHandler) Handle(bot *linebot.Client, event *linebot.Event) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Show")).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func(handler ShowToDoListHandler) fetchData(bot *linebot.Client, event *linebot.Event) {

}