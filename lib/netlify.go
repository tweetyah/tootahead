package lib

import (
	"database/sql"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/bmorrisondev/go-utils"
	"github.com/golang-jwt/jwt"
)

type NetlifyRouter struct {
	Get    func(events.APIGatewayProxyRequest, jwt.MapClaims, *sql.DB) (events.APIGatewayProxyResponse, error)
	Post   func(events.APIGatewayProxyRequest, jwt.MapClaims, *sql.DB) (events.APIGatewayProxyResponse, error)
	Put    func(events.APIGatewayProxyRequest, jwt.MapClaims, *sql.DB) (events.APIGatewayProxyResponse, error)
	Patch  func(events.APIGatewayProxyRequest, jwt.MapClaims, *sql.DB) (events.APIGatewayProxyResponse, error)
	Delete func(events.APIGatewayProxyRequest, jwt.MapClaims, *sql.DB) (events.APIGatewayProxyResponse, error)
}

func (nr *NetlifyRouter) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	authHeader := request.Headers["authorization"]
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	claims, isLoggedIn := ValidateToken(authHeader)
	if !isLoggedIn {
		return utils.UnauthorizedResponse(nil)
	}

	db, err := GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	if request.HTTPMethod == "GET" && nr.Get != nil {
		return nr.Get(request, claims, db)
	}
	if request.HTTPMethod == "POST" && nr.Post != nil {
		return nr.Post(request, claims, db)
	}
	if request.HTTPMethod == "PUT" && nr.Put != nil {
		return nr.Put(request, claims, db)
	}
	if request.HTTPMethod == "PATCH" && nr.Patch != nil {
		return nr.Patch(request, claims, db)
	}
	if request.HTTPMethod == "DELETE" && nr.Delete != nil {
		return nr.Delete(request, claims, db)
	}
	return utils.NotFoundResponse()
}
