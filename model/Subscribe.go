package model

import "github.com/go-sql-driver/mysql"

type Subscribe struct {
	Id string
	Name string
	DeletedAt mysql.NullTime
}
