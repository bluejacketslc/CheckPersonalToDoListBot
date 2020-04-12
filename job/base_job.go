package job

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"time"
)

var (
	Bot *linebot.Client
	currentLocation, _ = time.LoadLocation("Asia/Jakarta")
)

type BaseJob interface {
	Execute()
}