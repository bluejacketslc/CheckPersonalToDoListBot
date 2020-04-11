package model

import "github.com/go-sql-driver/mysql"

type ToDo struct {
	Id string
	UserId string
	Name string
	Deadline mysql.NullTime
	DeletedAt mysql.NullTime
}
