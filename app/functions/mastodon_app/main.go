package main

import (
	"database/sql"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bmorrisondev/go-utils"
	"github.com/golang-jwt/jwt"
	"github.com/tweetyah/lib"
)

func main() {
	router := lib.NetlifyRouter{
		Get: Get,
	}
	lambda.Start(router.Handler)
}

type RepsonseBody struct {
	ClientId string `json:"clientId"`
}

func Get(request events.APIGatewayProxyRequest, claims jwt.MapClaims, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	// lookup in the db if we have those tokens, return them if so
	// query := `select client_id from mastodon_apps where domain = ? limit 1`
	//
	// db, err := lib.GetDatabase()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// res := db.QueryRow(query, )

	// if not, register an app on that server, store the secrets, and return them
	return utils.OkResponse(nil)
}
