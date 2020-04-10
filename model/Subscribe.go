package model

import "time"

type Subscribe struct {
	Id string
	Name string
	DeletedAt *time.Time
}
