package main

import (
	"database/sql"
	"log"
	"os"

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

type ResponseBody struct {
	ClientId string `json:"clientId"`
}

func Get(request events.APIGatewayProxyRequest, claims jwt.MapClaims, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	domain := request.QueryStringParameters["domain"]
	requiresAppRegistration := false

	// lookup in the db if we have those tokens, return them if so
	query := `select client_id from mastodon_apps where domain = ? and redirect_uri = ? limit 1`

	db, err := lib.GetDatabase()
	if err != nil {
		log.Fatal(err)
	}
	res := db.QueryRow(query, domain, os.Getenv("VITE_REDIRECT_URI"))

	var response ResponseBody
	err = res.Scan(&response.ClientId)
	if err != nil && err == sql.ErrNoRows {
		requiresAppRegistration = true
	}
	// Another error was hit
	if err != nil && !requiresAppRegistration {
		log.Fatal(err)
	}

	// if not, register an app on that server, store the secrets, and return them
	if requiresAppRegistration {
		appRegsitration, err := lib.RegisterMastodonApp(domain)
	}

	return utils.OkResponse(nil)
}
