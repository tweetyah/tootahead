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
		rows := GetScheduledPosts()
		log.Printf("found %v tweets!", len(rows))
		users := RowsToUserTweets(rows)
		for _, user := range users {
			log.Printf("processing user %v", user.Id)
			if user.Service != nil && *user.Service == lib.AUTH_PROVIDER_TWITTER {
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

				// for _, p := range user.Posts {
				// 	HandleTweet(p, *user.AccessToken)
				// }
			}

			if user.Service != nil && *user.Service == lib.AUTH_PROVIDER_MASTODON {
				for _, p := range user.Posts {
					HandleMastodonPost(p, *user.MastodonDomain, *user.AccessToken)
				}
			}

			log.Printf("finished handling %v", user.Id)
		}
		log.Println("done!")
		time.Sleep(1 * time.Minute)
	}
}

type GetScheduledPostsDbResult struct {
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
	Service           *int
	MastodonDomain    *string
}

type User struct {
	Id                int64
	AccessToken       *string
	RefreshToken      *string
	AccessTokenExpiry *time.Time
	Posts             []PostRecord
	Service           *int
	MastodonDomain    *string
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

type PostRecord struct {
	Id           int64
	Text         *string
	SendAt       *time.Time
	RetweetAt    *time.Time
	IsThread     *bool
	ThreadCount  *int64
	ThreadParent *int64
	ThreadOrder  *int64
}

func RowsToUserTweets(rows []GetScheduledPostsDbResult) map[int64]User {
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
				Posts:             []PostRecord{},
				Service:           el.Service,
				MastodonDomain:    el.MastodonDomain,
			}
		}

		user := userMap[*el.UserId]
		user.Posts = append(userMap[*el.UserId].Posts, PostRecord{
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
func GetScheduledPosts() []GetScheduledPostsDbResult {
	query := `
		select
			p.id,
			p.text,
			p.send_at,
			p.retweet_at,
			p.is_thread,
			p.thread_count,
			p.id_user,
			p.thread_parent,
			p.thread_order,
			ut.access_token,
			ut.refresh_token,
			ut.access_token_expiry,
			p.service,
			ut.mastodon_domain
		from
			posts p
		left join
			user_tokens ut on p.id_user = ut.user_id
		where
			status = 0 and send_at < NOW()`
	db, err := GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	results, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var rows []GetScheduledPostsDbResult
	for results.Next() {
		var r GetScheduledPostsDbResult
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
			&r.Service,
			&r.MastodonDomain,
		)
		if err != nil {
			log.Fatal(err)
		}

		rows = append(rows, r)
	}

	return rows
}

// Sends a tweet and updates the db record
func HandleTweet(t PostRecord, accessToken string) {
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
		query := "update posts set id_sent = ?, status = 1 where id = ?"
		_, err = db.Exec(query, results.SentId, t.Id)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// TODO: Update this to forward error to Discord
		query := "update posts set status = 2, error = ? where id = ?"
		_, err = db.Exec(query, results.Status, t.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("Post %v processed successfully!", t.Id)
}

func HandleMastodonPost(p PostRecord, instanceDomain string, accessToken string) {
	db, err := GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	results, err := lib.SendMastodonPost(instanceDomain, *p.Text, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	if results.Error == nil {
		query := "update posts set id_sent = ?, status = 1 where id = ?"
		_, err = db.Exec(query, results.ID, p.Id)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// TODO: Update this to forward error to Discord
		query := "update posts set status = 2, error = ? where id = ?"
		_, err = db.Exec(query, *results.Error, p.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("Post %v processed successfully!", p.Id)

}
