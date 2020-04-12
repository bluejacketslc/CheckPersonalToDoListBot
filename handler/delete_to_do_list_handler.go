package handler

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"strings"
	"time"
	"todoreminder/helpers"
	"todoreminder/model"
)

var (
	deleteHelpInstructions =
		"Command Usage:\n" +
		"/remove [To Do ID]"
)

type DeleteToDoListHandler struct {}

func(handler DeleteToDoListHandler) Handle(bot *linebot.Client, event *linebot.Event) {
	dbConnection := helpers.CreateConnection()
	userId := event.Source.UserID
	toDoId := handler.fetchData(bot, event)
	selectedToDo := model.ToDo{
		Id:        toDoId,
		UserId:    userId,
		Name:      "",
		Deadline:  mysql.NullTime{},
		DeletedAt: mysql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	if handler.find(dbConnection, &selectedToDo) {
		handler.delete(dbConnection, &selectedToDo)

		_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("You have success delete a to do.")).Do()
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Your To Do List you want to delete is not found.")).Do()
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func(handler DeleteToDoListHandler) fetchData(bot *linebot.Client, event *linebot.Event) string {
	var toDoId string
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		rawTextMessage := message.Text
		arrSplitString := strings.SplitN(rawTextMessage, " ", 2)
		if len(arrSplitString) != 2 {
			_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(deleteHelpInstructions)).Do()
			if err != nil {
				log.Fatal(err.Error())
			}
			log.Fatal("Not Enough Arguments")
		}
		toDoId = arrSplitString[1]
	}

	return toDoId
}

func(handler DeleteToDoListHandler) find(dbConnection *sql.DB, t *model.ToDo) bool {
	query := "SELECT * FROM todo WHERE id=? AND user_id=? AND deleted_at IS NULL"
	currentStatement, err := dbConnection.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	results, err := currentStatement.Query(t.Id, t.UserId)
	if err != nil {
		log.Fatal(err.Error())
	}

	if results.Next() {
		return true
	}

	return false
}

func(handler DeleteToDoListHandler) delete(dbConnection *sql.DB, t *model.ToDo) {
	query := "UPDATE todo SET deleted_at=? WHERE id=? AND user_id=?"
	currentStatement, err := dbConnection.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = currentStatement.Exec(t.DeletedAt, t.Id, t.UserId)
	if err != nil {
		log.Fatal(err.Error())
	}
}