package lib

import (
	"database/sql"
	"log"
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

func SavePostToDb(userId string, serviceId int, post Post) (*Post, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, errors.Wrap(err, "(SavePostToDb) GetDatabase")
	}

	query := "insert into posts (text, send_at, retweet_at, id_user, service) values (?, ?, ?, ?, ?)"
	results, err := db.Exec(query, post.Text, post.GetSendAtSqlTimestamp(), post.GetResendAtSqlTimestamp(), userId, serviceId)
	if err != nil {
		return nil, errors.Wrap(err, "(SavePostToDb) db.Exec")
	}

	lastInserted, err := results.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "(SavePostToDb) result.LastInsertedId")
	}
	post.Id = &lastInserted

	log.Println("returning", post)

	return &post, nil
}

func SaveThreadToDb(userId string, serviceId int, posts []Post) (*Post, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, errors.Wrap(err, "(SaveThreadToDb) GetDatabase")
	}

	threadOrder := 1
	threadCount := len(posts)
	query := "insert into posts (text, is_thread, thread_order, thread_count, send_at, retweet_at, id_user, service) values (?, true, ?, ?, ?, ?, ?, ?)"
	threadStart := posts[0]
	results, err := db.Exec(query,
		threadStart.Text,
		threadOrder,
		threadCount,
		threadStart.GetSendAtSqlTimestamp(),
		threadStart.GetResendAtSqlTimestamp(),
		userId,
		serviceId,
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
	query = "insert into posts (text, is_thread, thread_order, thread_parent, send_at, retweet_at, id_user, service) values (?, true, ?, ?, ?, ?, ?, ?)"
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
		params = append(params, serviceId)
	}
	_, err = db.Exec(query, params)
	if err != nil {
		return nil, errors.Wrap(err, "(SavePostToDb) result.LastInsertedId")
	}

	return &threadStart, nil
}

func GetUserBySocialLogin(providerType int, providerId string) (*User, error) {
	query := `select u.id, u.last_login from users u 
		left outer join auth_providers ap on ap.user_id = u.id 
		where ap.type = ? and ap.service_id = ?
		limit 1`

	log.Println(query)

	db, err := GetDatabase()
	if err != nil {
		return nil, errors.Wrap(err, "(GetUserBySocialLogin) GetDatabase")
	}

	row := db.QueryRow(query, providerType, providerId)

	var record User
	err = row.Scan(&record.Id, &record.LastLogin)
	// User doesnt exist
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "(GetUserBySocialLogin) row.Scan")
	}

	return &record, nil
}

func CreateUserFromSocialLogin(providerType int, providerId string) (*User, error) {
	user := User{
		LastLogin: time.Now(),
	}

	db, err := GetDatabase()
	if err != nil {
		return nil, errors.Wrap(err, "(CreateUser) GetDatabase")
	}

	query := "insert into users (last_login) values (?)"
	res, err := db.Exec(query, SqlTimeStampFromTime(&user.LastLogin))
	if err != nil {
		return nil, errors.Wrap(err, "(CreatUser) insert into users")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "(CreateUser) last inserted id")
	}
	user.Id = &id

	query = "insert into auth_providers (user_id, type, service_id) values (?, ?, ?)"
	_, err = db.Exec(query, id, providerType, providerId)
	if err != nil {
		return nil, errors.Wrap(err, "(CreateUser) insert into auth_providers")
	}

	return &user, nil
}
