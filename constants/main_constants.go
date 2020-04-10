package constants

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"todoreminder/controller"
)

var EventListeners = map[linebot.EventType]interface{} {
	linebot.EventTypeMessage: (*controller.EventTypeMessageController).Execute,
	linebot.EventTypeFollow: (*controller.EventTypeFollowController).Execute,
}