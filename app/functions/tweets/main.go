package main

import (
	"database/sql"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	utils "github.com/bmorrisondev/go-utils"
	"github.com/golang-jwt/jwt"
	"github.com/tweetyah/lib"
)

func main() {
	router := lib.NetlifyRouter{
		Get:  Get,
		Post: Post,
	}
	lambda.Start(router.Handler)
}

func Get(request events.APIGatewayProxyRequest, claims jwt.MapClaims, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "hello world!",
	}, nil
}

func Post(request events.APIGatewayProxyRequest, claims jwt.MapClaims, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	var tweets []lib.Tweet
	err := json.Unmarshal([]byte(request.Body), &tweets)
	if err != nil {
		return utils.ErrorResponse(err, "json.Unmarshal")
	}

	userId := claims["twitter:user_id"]

	if len(tweets) == 1 {
		query := "insert into tweets (text, send_at, retweet_at, id_user) values (?, ?, ?, ?)"
		t := tweets[0]
		results, err := db.Exec(query, t.Text, t.GetSendAtSqlTimestamp(), t.GetRetweetAtSqlTimestamp(), userId)
		if err != nil {
			return utils.ErrorResponse(err, "db.Exec")
		}

		lastInserted, err := results.LastInsertId()
		if err != nil {
			return utils.ErrorResponse(err, "results.LastInsertedId")
		}

		t.Id = &lastInserted

		jstr, err := utils.ConvertToJsonString(t)
		if err != nil {
			return utils.ErrorResponse(err, "utils.ConvertToJsonString")
		}
		return utils.OkResponse(&jstr)
	} else {
		threadOrder := 1
		threadCount := len(tweets)
		query := "insert into tweets (text, is_thread, thread_order, thread_count, send_at, retweet_at, id_user) values (?, true, ?, ?, ?, ?, ?)"
		threadStart := tweets[0]
		results, err := db.Exec(query,
			threadStart.Text,
			threadOrder,
			threadCount,
			threadStart.GetSendAtSqlTimestamp(),
			threadStart.GetRetweetAtSqlTimestamp(),
			userId,
		)
		if err != nil {
			return utils.ErrorResponse(err, "(Post) db.Exec")
		}
		lastInserted, err := results.LastInsertId()
		if err != nil {
			return utils.ErrorResponse(err, "(Post) results.LastInsertedId")
		}
		threadStart.Id = &lastInserted
		threadStart.ThreadCount = &threadCount

		var params []interface{}
		query = "insert into tweets (text, is_thread, thread_order, thread_parent, send_at, retweet_at, id_user) values (?, true, ?, ?, ?, ?, ?)"
		for idx, el := range tweets {
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
			params = append(params, el.GetRetweetAtSqlTimestamp())
			params = append(params, userId)
		}
		_, err = db.Exec(query, params)
		if err != nil {
			return utils.ErrorResponse(err, "(Post) db.Exec on thread")
		}

		jstr, err := utils.ConvertToJsonString(threadStart)
		if err != nil {
			return utils.ErrorResponse(err, "(Post) utils.ConvertToJsonString")
		}
		return utils.OkResponse(&jstr)
	}
}
