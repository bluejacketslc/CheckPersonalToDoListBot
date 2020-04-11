package controller

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"strings"
	"todoreminder/handler"
)

//var commands = map[string] handler.BaseHandler {
//	"/subscribe": handler.SubscribeHandler{},
//	"/unsubscribe": handler.UnsubscribeHandler{},
//	"/add":       handler.AddToDoListHandler{},
//	"/list":      handler.ShowToDoListHandler{},
//	"/remove":    handler.DeleteToDoListHandler{},
//	"/help": handler.HelpToDoListHandler{},
//}

type EventTypeMessageController struct {}

func (controller EventTypeMessageController) Execute(bot *linebot.Client, event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		//commands[message.Text].Handle(bot, event)
		if strings.Contains(message.Text, "/subscribe") {
			handler.SubscribeHandler{}.Handle(bot, event)
		} else if strings.Contains(message.Text, "/unsubscribe") {
			handler.UnsubscribeHandler{}.Handle(bot, event)
		} else if strings.Contains(message.Text, "/add") {
			handler.AddToDoListHandler{}.Handle(bot, event)
		} else if strings.Contains(message.Text, "/list") {
			handler.ShowToDoListHandler{}.Handle(bot, event)
		} else if strings.Contains(message.Text, "/remove") {
			handler.DeleteToDoListHandler{}.Handle(bot, event)
		} else if strings.Contains(message.Text, "/help") {
			handler.HelpToDoListHandler{}.Handle(bot, event)
		}
	}
}