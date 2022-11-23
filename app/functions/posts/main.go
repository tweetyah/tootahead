package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

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

	log.Println(posts)

	userId := claims["user_id"].(string)
	service := claims["service_id"].(string)
	serviceId, err := strconv.Atoi(service)
	if err != nil {
		return utils.ErrorResponse(err, "(Post) cast service to num")
	}

	if len(posts) == 1 {
		updated, err := lib.SavePostToDb(userId, serviceId, posts[0])
		if err != nil {
			return utils.ErrorResponse(err, "(Post) save post to db")
		}

		jstr, err := utils.ConvertToJsonString(updated)
		if err != nil {
			return utils.ErrorResponse(err, "utils.ConvertToJsonString")
		}
		return utils.OkResponse(&jstr)
	} else {
		threadStart, err := lib.SaveThreadToDb(userId, serviceId, posts)
		if err != nil {
			return utils.ErrorResponse(err, "(Post) save thread to db")
		}

		jstr, err := utils.ConvertToJsonString(threadStart)
		if err != nil {
			return utils.ErrorResponse(err, "(Post) utils.ConvertToJsonString")
		}
		return utils.OkResponse(&jstr)
	}
}
