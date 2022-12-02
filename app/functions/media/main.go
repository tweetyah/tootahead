package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bmorrisondev/go-utils"
	"github.com/golang-jwt/jwt"
	"github.com/tweetyah/lib"
)

func main() {
	router := lib.NetlifyRouter{
		Post: Post,
	}
	lambda.Start(router.Handler)
}

func Post(request events.APIGatewayProxyRequest, claims jwt.MapClaims, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	var requestBody PostRequestBody
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return utils.ErrorResponse(err, "json.Unmarshal")
	}

	userId := claims["user_id"].(string)
	userIdNum, err := strconv.Atoi(userId)
	if err != nil {
		return utils.ErrorResponse(err, "(Post) failed to cast user id to int")
	}

	domain, token, err := lib.GetMastodonInstanceAndTokenByUser(userIdNum)
	if err != nil {
		return utils.ErrorResponse(err, "(Post) get instance and token by user")
	}

	log.Println(*token)

	res, err := lib.UploadMediaToMastodon(*domain, *token, requestBody.File)
	if err != nil {
		return utils.ErrorResponse(err, "(Post) upload media to mastodon")
	}

	jstr, err := utils.ConvertToJsonString(res)
	if err != nil {
		return utils.ErrorResponse(err, "(Post) marshal")
	}

	return utils.OkResponse(&jstr)
}

type PostRequestBody struct {
	File string `json:"file"`
}
