package handler

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"time"
	"todoreminder/helpers"
	"todoreminder/model"
)

type UnsubscribeHandler struct {}

func(handler UnsubscribeHandler) Handle(bot *linebot.Client, event *linebot.Event) {
	dbConnection := helpers.CreateConnection()
	userId := event.Source.UserID

	currentSubscriber := handler.find(dbConnection, userId)
	if currentSubscriber != nil || currentSubscriber.DeletedAt.Valid == false {
		currentSubscriber.DeletedAt.Time = time.Now()
		currentSubscriber.DeletedAt.Valid = true

		handler.update(dbConnection, *currentSubscriber)
		_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("You have set daily to-do-list reminder to off. Reminder will not be shown. Enter \"/subscribe\" to activate reminder again.")).Do()
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("You have not subscribed yet.")).Do()
		if err != nil {
			log.Fatal(err.Error())
		}
	}


	err := dbConnection.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (handler UnsubscribeHandler) find(dbConnection *sql.DB, userId string) *model.Subscribe {
	query := "SELECT * FROM subscribers WHERE id=?"
	currentStatement, err := dbConnection.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	results, err := currentStatement.Query(userId)
	if results.Next() {
		var id string
		var name string
		var deletedAt mysql.NullTime
		err = results.Scan(&id, &name, &deletedAt)
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

func (handler UnsubscribeHandler) update(dbConnection *sql.DB, s model.Subscribe) {
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
