package handler

import "github.com/line/line-bot-sdk-go/linebot"

type BaseHandler interface {
	Handle(bot *linebot.Client, event *linebot.Event)
}
