package constants

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"todoreminder/controller"
)

var EventListeners = map[linebot.EventType] controller.BaseController {
	linebot.EventTypeMessage: controller.EventTypeMessageController{},
	linebot.EventTypeFollow: controller.EventTypeFollowController{},
}