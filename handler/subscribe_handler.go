package handler

import (
	"database/sql"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"todoreminder/helpers"
	"todoreminder/model"
)

type SubscribeHandler struct {}

func (handler SubscribeHandler) Handle(bot *linebot.Client, event *linebot.Event) {
	dbConnection := helpers.CreateConnection()
	userId := event.Source.UserID
	userName, err := bot.GetProfile(userId).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
	currentSubscriber := handler.find(dbConnection, userId)
	if currentSubscriber != nil {
		currentSubscriber.DeletedAt = ""

		handler.update(dbConnection, *currentSubscriber)
	} else {
		newSubscriber := model.Subscribe{
			Id:        userId,
			Name:      userName.DisplayName,
			DeletedAt: "",
		}

		handler.create(dbConnection, newSubscriber)
	}

	_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("You have set daily to-do-list reminder on. Reminder will send message every 07.00 GMT +7")).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = dbConnection.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (handler SubscribeHandler) find(dbConnection *sql.DB, userId string) *model.Subscribe {
	query := "SELECT * FROM subscribers WHERE id=?"
	currentStatement, err := dbConnection.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	results, err := currentStatement.Query(userId)
	if results.Next() {
		var name, deletedAt string
		err = results.Scan(&name, &deletedAt)
		if err != nil {
			log.Fatal(err.Error())
		}
		return &model.Subscribe{
			Id:        userId,
			Name:      name,
			DeletedAt: deletedAt,
		}
	}

	return nil
}

func (handler SubscribeHandler) update(dbConnection *sql.DB, s model.Subscribe) {
	query := "UPDATE subscribers SET deleted_at=?"
	currentStatement, err := dbConnection.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = currentStatement.Exec(s.DeletedAt)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (handler SubscribeHandler) create(dbConnection *sql.DB, s model.Subscribe) {
	query := "INSERT INTO subscribers (id, name) VALUES ($1, $2)"
	currentStatement, err := dbConnection.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = currentStatement.Exec(s.Id, s.Name)
	if err != nil {
		log.Fatal(err.Error())
	}
}
