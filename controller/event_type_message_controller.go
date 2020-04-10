package controller

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
)

type EventTypeMessageController struct {}

func (controller EventTypeMessageController) Execute(bot *linebot.Client, event *linebot.Event) {
	fmt.Println("EventTypeMessage")
}