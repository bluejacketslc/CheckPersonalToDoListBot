package controller

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
)

type EventTypeFollowController struct {}

func (controller *EventTypeFollowController) Execute(event *linebot.Event) {
	fmt.Println("EventTypeFollow")
}