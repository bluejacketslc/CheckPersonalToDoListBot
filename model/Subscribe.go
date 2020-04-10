package model

import (
	"database/sql/driver"
	"time"
)

type Subscribe struct {
	Id string
	Name string
	DeletedAt *time.Time
}

// Scan implements the Scanner interface.
func (nt *Subscribe) Scan(value interface{}) error {
	var deletedAt time.Time
	deletedAt = value.(time.Time)
	nt.DeletedAt = &deletedAt
	return nil
}

// Value implements the driver Valuer interface.
func (nt Subscribe) Value() (driver.Value, error) {
	return nt.DeletedAt, nil
}
