package controller

import "github.com/line/line-bot-sdk-go/linebot"

type BaseController interface {
	Execute(bot *linebot.Client, event *linebot.Event)
}