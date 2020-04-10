package controller

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"todoreminder/handler"
)

var commands = map[string] handler.BaseHandler {
	"/subscribe": handler.SubscribeController{},
	"/add":       handler.AddToDoListController{},
	"/list":      handler.ShowToDoListController{},
	"/remove":    handler.DeleteToDoListController{},
	"/help": handler.HelpToDoListHandler{},
}

type EventTypeMessageController struct {}

func (controller EventTypeMessageController) Execute(bot *linebot.Client, event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		commands[message.Text].Handle(bot, event)
	}
}