package helpers

import (
	"database/sql"
	"log"
	"os"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/")
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
