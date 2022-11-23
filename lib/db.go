package lib

import (
	"time"

	"github.com/pkg/errors"
)

func SaveTwitterAccessToken(userId int64, accessToken string, tokenExpiration time.Time, refreshToken string) error {
	query := `insert into user_tokens
			(id, user_id, access_token, access_token_expiry, refresh_token) values (?, ?, ?, ?, ?)
		on duplicate key update
			access_token = ?, access_token_expiry = ?, refresh_token = ?`
	db, err := GetDatabase()
	if err != nil {
		return errors.Wrap(err, "(SaveTwitterAccessToken) lib.GetDatabase")
	}
	_, err = db.Exec(
		query,
		userId,
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

func SavePostToDb(userId string, post Post) (*Post, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, errors.Wrap(err, "(SavePostToDb) GetDatabase")
	}

	query := "insert into posts (text, send_at, retweet_at, id_user) values (?, ?, ?, ?)"
	results, err := db.Exec(query, post.Text, post.GetSendAtSqlTimestamp(), post.GetResendAtSqlTimestamp(), userId)
	if err != nil {
		return nil, errors.Wrap(err, "(SavePostToDb) db.Exec")
	}

	lastInserted, err := results.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "(SavePostToDb) result.LastInsertedId")
	}

	post.Id = &lastInserted
	return &post, nil
}

func SaveThreadToDb(userId string, posts []Post) (*Post, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, errors.Wrap(err, "(SaveThreadToDb) GetDatabase")
	}

	threadOrder := 1
	threadCount := len(posts)
	query := "insert into posts (text, is_thread, thread_order, thread_count, send_at, retweet_at, id_user) values (?, true, ?, ?, ?, ?, ?)"
	threadStart := posts[0]
	results, err := db.Exec(query,
		threadStart.Text,
		threadOrder,
		threadCount,
		threadStart.GetSendAtSqlTimestamp(),
		threadStart.GetResendAtSqlTimestamp(),
		userId,
	)
	if err != nil {
		return nil, errors.Wrap(err, "(SaveThreadToDb) db.Exec threadstart")
	}
	lastInserted, err := results.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "(SaveThreadToDb) result.LastInsertedId")
	}
	threadStart.Id = &lastInserted
	threadStart.ThreadCount = &threadCount

	var params []interface{}
	query = "insert into posts (text, is_thread, thread_order, thread_parent, send_at, retweet_at, id_user) values (?, true, ?, ?, ?, ?, ?)"
	for idx, el := range posts {
		// Skip the first tweet since it was inserted earlier
		if idx == 0 {
			continue
		}
		threadOrder++
		if idx > 1 {
			query += ",(?, true, ?, ?, ?, ?)"
		}
		params = append(params, el.Text)
		params = append(params, threadOrder)
		params = append(params, threadStart.Id)
		params = append(params, el.GetSendAtSqlTimestamp())
		params = append(params, el.GetResendAtSqlTimestamp())
		params = append(params, userId)
	}
	_, err = db.Exec(query, params)
	if err != nil {
		return nil, errors.Wrap(err, "(SavePostToDb) result.LastInsertedId")
	}

	return &threadStart, nil
}
