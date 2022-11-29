package lib

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var _db *sql.DB

func GetDatabase() (*sql.DB, error) {
	if _db == nil {
		db, err := sql.Open("mysql", os.Getenv("DSN"))
		if err != nil {
			return nil, errors.Wrap(err, "(GetDatabase) failed to open")
		}
		_db = db
	}
	return _db, nil
}

func SqlTimeStampFromTime(t *time.Time) *string {
	if t != nil {
		v := t.Format("2006-01-02 15:04:05")
		return &v
	}
	return nil
}
