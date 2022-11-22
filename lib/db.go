package lib

import (
	"time"

	"github.com/pkg/errors"
)

func SaveTwitterAccessToken(userId int64, accessToken string, tokenExpiration time.Time, refreshToken string) error {
	query := `insert into users
			(id, access_token, access_token_expiry, refresh_token) values (?, ?, ?, ?)
		on duplicate key update
			access_token = ?, access_token_expiry = ?, refresh_token = ?`
	db, err := GetDatabase()
	if err != nil {
		return errors.Wrap(err, "(SaveTwitterAccessToken) lib.GetDatabase")
	}
	_, err = db.Exec(
		query,
		userId,
		accessToken,
		tokenExpiration,
		refreshToken,
		accessToken,
		tokenExpiration,
		refreshToken,
	)
	if err != nil {
		return errors.Wrap(err, "(SaveTwitterAccessToken) db.Exec")
	}
	return nil
}
