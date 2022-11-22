package lib

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	return db, err
}
