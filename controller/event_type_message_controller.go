package controller

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"todoreminder/handler"
)

var commands = map[string] handler.BaseHandler {
	"/subscribe": handler.SubscribeHandler{},
	"/unsubscribe": handler.UnsubscribeHandler{},
	"/add":       handler.AddToDoListHandler{},
	"/list":      handler.ShowToDoListHandler{},
	"/remove":    handler.DeleteToDoListHandler{},
	"/help": handler.HelpToDoListHandler{},
}

type EventTypeMessageController struct {}

func (controller EventTypeMessageController) Execute(bot *linebot.Client, event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		commands[message.Text].Handle(bot, event)
	}
}