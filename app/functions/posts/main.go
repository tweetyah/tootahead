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
	var posts []lib.Post
	err := json.Unmarshal([]byte(request.Body), &posts)
	if err != nil {
		return utils.ErrorResponse(err, "json.Unmarshal")
	}

	userId := claims["user_id"].(string)

	if len(posts) == 1 {
		updated, err := lib.SavePostToDb(userId, posts[0])
		jstr, err := utils.ConvertToJsonString(updated)
		if err != nil {
			return utils.ErrorResponse(err, "utils.ConvertToJsonString")
		}
		return utils.OkResponse(&jstr)
	} else {
		threadStart, err := lib.SaveThreadToDb(userId, posts)
		jstr, err := utils.ConvertToJsonString(threadStart)
		if err != nil {
			return utils.ErrorResponse(err, "(Post) utils.ConvertToJsonString")
		}
		return utils.OkResponse(&jstr)
	}
}
