package controller

import "fmt"

type EventTypeMessageController struct {}

func (controller *EventTypeMessageController) Execute() {
	fmt.Println("EventTypeMessage")
}