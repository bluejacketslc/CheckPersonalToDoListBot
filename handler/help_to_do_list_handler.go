package handler

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

var (
	helpInstructions =
		"To Do Reminder\n" +
		"==============\n" +
		"Command Format: [options]\n" +
		"Options:\n" +
		"/subscribe -> Activate Daily Reminder every 07.00 GMT +7\n" +
		"/unsubscribe -> Deactivate Daily Reminder\n" +
		"/add [deadline in yyyy-mm-dd] [task name] -> Add new To Do List\n" +
		"/list -> Show 5 near deadline To Do Lists\n" +
		"/list [date in yyyy-mm-dd] -> Show all To Do Lists on that date\n" +
		"/remove [To Do ID] -> Remove To Do List based on ID\n" +
		"/help -> Show Usage of Commands\n" +
		"\n" +
		"If you want to add any ideas to this code, please contact me.\n" +
		"Developer Email: christopherlimawan@gmail.com\n" +
		"Hardwork and Extraordinary Effort Makes Success Comes Closer ~ Bluejack 18-1"
)

type HelpToDoListHandler struct {}

func (handler HelpToDoListHandler) Handle(bot *linebot.Client, event *linebot.Event) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(helpInstructions)).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}
