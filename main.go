package main

import (
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/robfig/cron"
	"log"
	"net/http"
	"os"
	"time"
	"todoreminder/controller"
	"todoreminder/job"
)

var EventListeners = map[linebot.EventType] controller.BaseController {
	linebot.EventTypeMessage: controller.EventTypeMessageController{},
	linebot.EventTypeFollow: controller.EventTypeFollowController{},
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func initBot() *linebot.Client {
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err.Error())
	}

	return bot
}

func setListeners(bot *linebot.Client, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, event := range events {
		EventListeners[event.Type].Execute(bot, event)
	}
}

func loadJobs(bot *linebot.Client, cronInstance *cron.Cron) {
	job.Bot = bot
	cronInstance.AddFunc("CRON_TZ=Asia/Jakarta 0 7 * * *", job.ReminderToDoListJob{}.Execute)
}

func main() {
	loadEnv()
	bot := initBot()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		setListeners(bot, r)
	})

	err := http.ListenAndServe(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"), nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	location, _ := time.LoadLocation("Asia/Jakarta")
	cronInstance := cron.New(cron.WithLocation(location))
	loadJobs(bot, cronInstance)
	cronInstance.Start()
}