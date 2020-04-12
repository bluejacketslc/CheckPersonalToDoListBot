package controller

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

var (
	welcomingMessage =
		"Check To Do List Bot\n" +
		"====================\n" +
		"Created by CP18-1\n" +
		"\n" +
		"Hi, %s.\n" +
		"This is a LINE bot that will help you for take a note and remind you for every task that you have.\n" +
		"For basic commands, why don't you try \"/help\" first?\n" +
		"Enjoy :D\n" +
		"\n" +
		"Created with love,\n" +
		"Christopher Limawan\n" +
		"Bluejack 18-1"
)

type EventTypeFollowController struct {}

func (controller EventTypeFollowController) Execute(bot *linebot.Client, event *linebot.Event) {
	profile, err := bot.GetProfile(event.Source.UserID).Do()
	welcomingMessage = fmt.Sprintf(welcomingMessage, profile.DisplayName)
	_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(welcomingMessage)).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}