package lib

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	return db, err
}

func SqlTimeStampFromTime(t *time.Time) *string {
	if t != nil {
		v := t.Format("2006-01-02 15:04:05")
		return &v
	}
	return nil
}
