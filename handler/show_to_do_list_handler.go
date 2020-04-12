package handler

import (
	"database/sql"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"strings"
	"time"
	"todoreminder/helpers"
	"todoreminder/model"
)

var (
	listHelpInstructions =
		"Command Usage:\n" +
		"/list -> Show 5 near deadline To Do Lists\n" +
		"/list [date in yyyy-mm-dd] -> Show all To Do Lists on that date"
)

type ShowToDoListHandler struct {}

func(handler ShowToDoListHandler) Handle(bot *linebot.Client, event *linebot.Event) {
	dbConnection = helpers.CreateConnection()
	argumentsLength := handler.getArgumentsLength(event)

	if argumentsLength == 1 {
		collectedToDos := handler.getNearDeadlineTodo(dbConnection, event.Source.UserID)
		handler.showDeadlineToDo(bot, event, collectedToDos)
	} else if argumentsLength == 2 {
		selectedDate := handler.fetchData(bot, event)
		collectedToDos := handler.getSelectedDeadlineToDo(dbConnection, userId, selectedDate)
		handler.showDeadlineToDo(bot, event, collectedToDos)
	} else {
		_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(listHelpInstructions)).Do()
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	err := dbConnection.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func(handler ShowToDoListHandler) getArgumentsLength(event *linebot.Event) int {
	var arguments = 0
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		rawMessage := message.Text
		arrSplitString := strings.Split(rawMessage, " ")
		arguments = len(arrSplitString)
	}

	return arguments
}

func(handler ShowToDoListHandler) getNearDeadlineTodo(dbConnection *sql.DB, userId string) []model.ToDo {
	var collectedToDos []model.ToDo
	query := "SELECT id, name, deadline FROM todo WHERE user_id=? AND deleted_at IS NULL AND deadline >= CAST(? AS DATE) ORDER BY deadline ASC LIMIT 8"
	currentStatement, err := dbConnection.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	currentDate := time.Now().In(currentLocation)
	results, err := currentStatement.Query(userId, currentDate.Format("2020-01-02"))

	defer results.Close()
	for results.Next() {
		currentToDo := model.ToDo{}
		err := results.Scan(&currentToDo.Id, &currentToDo.Name, &currentToDo.Deadline)
		if err != nil {
			log.Fatal(err.Error())
		}
		collectedToDos = append(collectedToDos, currentToDo)
	}

	return collectedToDos
}

func(handler ShowToDoListHandler) getSelectedDeadlineToDo(dbConnection *sql.DB, userId string, selectedDate string) []model.ToDo {
	var collectedToDos []model.ToDo
	query := "SELECT id, name, deadline FROM todo WHERE user_id=? AND deleted_at IS NULL AND deadline=CAST(? AS DATE)"
	currentStatement, err := dbConnection.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	// malicious date format will not be processed
	currentTime, err := time.Parse("2020-01-02", selectedDate)
	if err != nil {
		log.Fatal(err.Error())
	}
	results, err := currentStatement.Query(userId, currentTime.Format("2020-01-02"))
	defer results.Close()
	for results.Next() {
		currentToDo := model.ToDo{}
		err := results.Scan(&currentToDo.Id, &currentToDo.Name, &currentToDo.Deadline)
		if err != nil {
			log.Fatal(err.Error())
		}
		collectedToDos = append(collectedToDos, currentToDo)
	}

	return collectedToDos
}

func(handler ShowToDoListHandler) showDeadlineToDo(bot *linebot.Client, event *linebot.Event, collectedToDos []model.ToDo) {
	var currentMessage =
		"Current Near Deadline To Do Lists:\n"
	var lastMarkedDate = ""
	for _, currentToDo := range collectedToDos {
		currentMarkedDate := currentToDo.Deadline.Time.Format("2020-01-02")
		if currentMarkedDate != lastMarkedDate {
			if lastMarkedDate != "" {
				currentMessage += "\n"
			}
			currentMessage += currentMarkedDate + "\n"
			lastMarkedDate = currentMarkedDate
		}
		currentMessage += "(#" + currentToDo.Id + ") " + currentToDo.Name + "\n"
	}

	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(currentMessage)).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func(handler ShowToDoListHandler) fetchData(bot *linebot.Client, event *linebot.Event) string {
	var selectedDate string
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		rawTextMessage := message.Text
		arrSplitString := strings.Split(rawTextMessage, " ")

		selectedDate = arrSplitString[1]
	}

	return selectedDate
}