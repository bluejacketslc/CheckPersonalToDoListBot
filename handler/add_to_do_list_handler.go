package handler

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"strings"
	"time"
	"todoreminder/helpers"
	"todoreminder/model"
)

var (
	addHelpInstructions =
		"Command Usage:\n" +
		"/add [deadline in yyyy-mm-dd] [task name]"
)

type AddToDoListHandler struct {}

func(handler AddToDoListHandler) Handle(bot *linebot.Client, event *linebot.Event) {
	dbConnection := helpers.CreateConnection()
	generatedId := uuid.New().String()[:8]
	userId := event.Source.UserID
	name, rawDeadline := handler.fetchData(bot, event)
	deadline, err := time.Parse("2006-01-02", rawDeadline)
	if err != nil {
		log.Fatal(err.Error())
	}

	handler.create(dbConnection, &model.ToDo{
		Id:        generatedId,
		UserId:    userId,
		Name:      name,
		Deadline:  mysql.NullTime{
			Time:  deadline,
			Valid: true,
		},
		DeletedAt: mysql.NullTime{},
	})

	_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("New To Do List has been added with following details:\n" +
		"Task Name: " + name + "\nDeadline: " + deadline.Format("02 January 2006"))).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func(handler AddToDoListHandler) fetchData(bot *linebot.Client, event *linebot.Event) (string, string) {
	// format: /add [deadline] [name]
	var name, deadline string
	switch message := event.Message.(type){
	case *linebot.TextMessage:
		currentMessage := message.Text
		arrSplitString := strings.SplitN(currentMessage, " ", 3)
		if len(arrSplitString) != 3 {
			_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(addHelpInstructions)).Do()
			if err != nil {
				log.Fatal(err.Error())
			}
			log.Fatal("Not Enough Arguments")
		} else {
			deadline = arrSplitString[1]
			name = arrSplitString[2]
		}
	}

	return name, deadline
}

func(handler AddToDoListHandler) create(dbConnection *sql.DB, t *model.ToDo) {
	query := "INSERT INTO todo(id, user_id, name, deadline) VALUES (?,?,?,?)"
	currentStatement, err := dbConnection.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = currentStatement.Exec(t.Id, t.UserId, t.Name, t.Deadline)
	if err != nil {
		log.Fatal(err.Error())
	}
}