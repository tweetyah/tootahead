package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/tweetyah/lib"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	return db, err
}

func main() {
	// Get tweets
	godotenv.Load()

	for {
		log.Println("checking for tweets")
		rows := GetScheduledTweets()
		log.Printf("found %v tweets!", len(rows))
		users := RowsToUserTweets(rows)
		for _, user := range users {
			log.Printf("processing user %v", user.Id)
			didUpdate, err := user.ValidateTokens()
			if err != nil {
				log.Fatal(err)
			}
			if didUpdate {
				err = lib.SaveTwitterAccessToken(user.Id, *user.AccessToken, *user.AccessTokenExpiry, *user.RefreshToken)
				if err != nil {
					log.Fatal(err)
				}
			}

			for _, t := range user.Tweets {
				HandleTweet(t, *user.AccessToken)
			}

			log.Printf("finished handling %v", user.Id)
		}
		log.Println("done!")
		time.Sleep(1 * time.Minute)
	}
}

type GetScheduledTweetsDbResult struct {
	Id                int64
	Text              *string
	SendAt            *time.Time
	RetweetAt         *time.Time
	IsThread          *bool
	ThreadCount       *int64
	UserId            *int64
	ThreadParent      *int64
	ThreadOrder       *int64
	AccessToken       *string
	RefreshToken      *string
	AccessTokenExpiry *time.Time
}

type User struct {
	Id                int64
	AccessToken       *string
	RefreshToken      *string
	AccessTokenExpiry *time.Time
	Tweets            []TweetRecord
}

func (u *User) ValidateTokens() (bool, error) {
	// TODO: Check if the refresh token exists
	if u.AccessTokenExpiry == nil || (*u.AccessTokenExpiry).Before(time.Now()) {
		authResponse, err := lib.GetTwitterTokensViaRefresh(*u.RefreshToken)
		if err != nil {
			return false, errors.Wrap(err, "(ValidateTokens) lib.GetTwitterTokenViaRefresh")
		}
		accessTokenExpiry := time.Now().Add(time.Duration(authResponse.ExpiresIn-60) * time.Second)

		u.AccessToken = &authResponse.AccessToken
		u.RefreshToken = &authResponse.RefreshToken
		u.AccessTokenExpiry = &accessTokenExpiry
		return true, nil
	}
	return false, nil
}

type TweetRecord struct {
	Id           int64
	Text         *string
	SendAt       *time.Time
	RetweetAt    *time.Time
	IsThread     *bool
	ThreadCount  *int64
	ThreadParent *int64
	ThreadOrder  *int64
}

func RowsToUserTweets(rows []GetScheduledTweetsDbResult) map[int64]User {
	userMap := map[int64]User{}
	for _, el := range rows {
		if el.UserId == nil {
			continue
		}

		if _, ok := userMap[*el.UserId]; !ok {
			userMap[*el.UserId] = User{
				Id:                *el.UserId,
				AccessToken:       el.AccessToken,
				RefreshToken:      el.RefreshToken,
				AccessTokenExpiry: el.AccessTokenExpiry,
				Tweets:            []TweetRecord{},
			}
		}

		user := userMap[*el.UserId]
		user.Tweets = append(userMap[*el.UserId].Tweets, TweetRecord{
			Id:           el.Id,
			Text:         el.Text,
			SendAt:       el.SendAt,
			RetweetAt:    el.RetweetAt,
			IsThread:     el.IsThread,
			ThreadCount:  el.ThreadCount,
			ThreadParent: el.ThreadParent,
			ThreadOrder:  el.ThreadOrder,
		})
		userMap[*el.UserId] = user
	}
	return userMap
}

// Queries the database for all tweets to be sent
func GetScheduledTweets() []GetScheduledTweetsDbResult {
	query := `
		select
			t.id,
			t.text,
			t.send_at,
			t.retweet_at,
			t.is_thread,
			t.thread_count,
			t.id_user,
			t.thread_parent,
			t.thread_order,
			u.access_token,
			u.refresh_token,
			u.access_token_expiry
		from
			tweets t
		left join
			users u on t.id_user = u.id
		where
			send_at < NOW() and id_sent is null and status = 0`
	db, err := GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	results, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var rows []GetScheduledTweetsDbResult
	for results.Next() {
		var r GetScheduledTweetsDbResult
		err := results.Scan(
			&r.Id,
			&r.Text,
			&r.SendAt,
			&r.RetweetAt,
			&r.IsThread,
			&r.ThreadCount,
			&r.UserId,
			&r.ThreadParent,
			&r.ThreadOrder,
			&r.AccessToken,
			&r.RefreshToken,
			&r.AccessTokenExpiry,
		)
		if err != nil {
			log.Fatal(err)
		}

		rows = append(rows, r)
	}

	return rows
}

// Sends a tweet and updates the db record
func HandleTweet(t TweetRecord, accessToken string) {
	log.Printf("processing tweet %v", t.Id)
	db, err := GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	results, err := lib.SendTweet(*t.Text, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	if results.IsSuccess {
		query := "update tweets set id_sent = ?, status = 1 where id = ?"
		_, err = db.Exec(query, results.SentId, t.Id)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// TODO: Update this to forward error to Discord
		query := "update tweets set status = 2, error = ? where id = ?"
		_, err = db.Exec(query, results.Status, t.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("Tweet %v processed successfully!", t.Id)
}
