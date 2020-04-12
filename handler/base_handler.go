package handler

import (
	"database/sql"
	"github.com/line/line-bot-sdk-go/linebot"
	"time"
)

var (
	dbConnection *sql.DB
	currentLocation, _ = time.LoadLocation("Asia/Jakarta")
)

type BaseHandler interface {
	Handle(bot *linebot.Client, event *linebot.Event)
}
