package job

import (
	"database/sql"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"time"
	"todoreminder/helpers"
	"todoreminder/model"
)

type ReminderToDoListJob struct {}

func(job ReminderToDoListJob) Execute() {
	dbConnection := helpers.CreateConnection()

	defer dbConnection.Close()
	collectionSubscribers := job.getAllSubscribers(dbConnection)
	for _, subscriber := range collectionSubscribers {
		collectionToDos := job.getNearDeadlineTodo(dbConnection, subscriber.Id)
		job.sendMessage(subscriber.Id, collectionToDos)
	}
}

func(job ReminderToDoListJob) getAllSubscribers(dbConnection *sql.DB) []model.Subscribe {
	var collectionSubscribers []model.Subscribe
	query := "SELECT id FROM subscribers WHERE deleted_at IS NULL"
	results, err := dbConnection.Query(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer results.Close()
	for results.Next() {
		currentSubscriber := model.Subscribe{}
		err := results.Scan(&currentSubscriber.Id)
		if err != nil {
			log.Fatal(err.Error())
		}
		collectionSubscribers = append(collectionSubscribers, currentSubscriber)
	}

	return collectionSubscribers
}

func(job ReminderToDoListJob) getNearDeadlineTodo(dbConnection *sql.DB, userId string) []model.ToDo {
	var collectedToDos []model.ToDo
	query := "SELECT id, name, deadline FROM todo WHERE user_id=? AND deleted_at IS NULL AND deadline >= CAST(? AS DATE) ORDER BY deadline ASC LIMIT 8"
	currentStatement, err := dbConnection.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	currentDate := time.Now().In(currentLocation)
	results, err := currentStatement.Query(userId, currentDate.Format("2006-01-02"))

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

func(job ReminderToDoListJob) sendMessage(userId string, collectionToDos []model.ToDo) {
	profile, err := Bot.GetProfile(userId).Do()
	if err != nil {
		log.Fatal(err.Error())
	}

	var currentMessage =
		"Good Morning, " + profile.DisplayName + ".\n" +
		"Here's your updated current Near Deadline To Do Lists:\n"
	var lastMarkedDate = ""
	for _, currentToDo := range collectionToDos {
		currentMarkedDate := currentToDo.Deadline.Time.Format("2006-01-02")
		if currentMarkedDate != lastMarkedDate {
			if lastMarkedDate != "" {
				currentMessage += "\n"
			}
			currentMessage += currentMarkedDate + "\n"
			lastMarkedDate = currentMarkedDate
		}
		currentMessage += "(#" + currentToDo.Id + ") " + currentToDo.Name + "\n"
	}

	_, err = Bot.PushMessage(userId, linebot.NewTextMessage(currentMessage)).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}
