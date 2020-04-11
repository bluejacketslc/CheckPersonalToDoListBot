package helpers

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
