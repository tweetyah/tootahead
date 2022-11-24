package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	utils "github.com/bmorrisondev/go-utils"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/tweetyah/lib"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "POST" {
		res, err := Post(request)
		return &res, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 404,
	}, nil
}

type RequestBody struct {
	State string `json:"state"`
	Code  string `json:"code"`
}

type ResponseBody struct {
	AccessToken     string `json:"access_token"`
	Id              string `json:"id"`
	Name            string `json:"name"`
	ProfileImageUrl string `json:"profile_image_url"`
	Username        string `json:"username"`
	Service         string `json:"service"`
}

func Post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body RequestBody
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return utils.ErrorResponse(err, "json.Unmarshal")
	}

	var responseBody ResponseBody
	if body.State == "twitter" {
		rb, err := BuildTwitterResponse(body.Code)
		if err != nil {
			return utils.ErrorResponse(err, "(Post) BuildTwitterResponse")
		}
		responseBody = *rb
	} else {
		rb, err := BuildMastodonResponse(body.State, body.Code)
		if err != nil {
			return utils.ErrorResponse(err, "(Post) BuildMastodonResponse")
		}
		responseBody = *rb
	}

	jstr, err := utils.ConvertToJsonString(responseBody)
	if err != nil {
		return utils.ErrorResponse(err, "(Post) utils.ConvertToJsonString")
	}

	return utils.OkResponse(&jstr)
}

func main() {
	lambda.Start(handler)
}

func BuildTwitterResponse(code string) (*ResponseBody, error) {
	twitterAuthResp, err := lib.GetTwitterTokens(code)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildTwitterResponse) GetTwitterTokens")
	}

	userDetails, err := lib.GetTwitterUserDetails(twitterAuthResp.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildTwitterResponse) GetTwitterUserDetails")
	}

	idNum, err := strconv.Atoi(userDetails.Data.Id)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildTwitterResponse) convert user id to int")
	}
	authTokenExpiration := time.Now().Add(time.Duration(twitterAuthResp.ExpiresIn-60) * time.Second)
	err = lib.SaveTwitterAccessToken(int64(idNum), twitterAuthResp.AccessToken, authTokenExpiration, twitterAuthResp.RefreshToken)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildTwitterResponse) SaveTwitterAccessToken")
	}

	user, err := lib.GetUserBySocialLogin(1, userDetails.Data.Id)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildTwitterResponse) GetUserBySocialLogin")
	}
	if user == nil {
		user, err = lib.CreateUserFromSocialLogin(lib.AUTH_PROVIDER_TWITTER, userDetails.Data.Id)
		if err != nil {
			return nil, errors.Wrap(err, "(BuildTwitterResponse) CreateUserFromSocialLogin")
		}
	}

	tokenStr, err := MintJwt(*user.Id, lib.AUTH_PROVIDER_TWITTER)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildTwitterResponse) mint jwt")
	}

	rv := ResponseBody{
		AccessToken:     *tokenStr,
		Id:              fmt.Sprint(*user.Id),
		Name:            userDetails.Data.Name,
		ProfileImageUrl: userDetails.Data.ProfileImageUrl,
		Username:        userDetails.Data.Username,
		Service:         "twitter",
	}
	return &rv, nil
}

func BuildMastodonResponse(instanceDomain, code string) (*ResponseBody, error) {

	// TODO: Move this into lib
	var clientId string
	var clientSecret string
	query := "select client_id, client_secret from mastodon_apps where domain = ? and redirect_uri = ? limit 1"
	db, err := lib.GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	res := db.QueryRow(query, instanceDomain, os.Getenv("VITE_REDIRECT_URI"))
	err = res.Scan(&clientId, &clientSecret)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: end ===

	tokens, err := lib.GetMastodonTokens(instanceDomain, code, clientId, clientSecret)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildMastodonResponse) GetMastodonTokens")
	}

	userDetails, err := lib.GetMastodonUserDetails(instanceDomain, tokens.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildMastodonResponse) GetMastodonUserDetails")
	}

	user, err := lib.GetUserBySocialLogin(1, userDetails.ID)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildMastodonResponse) GetUserBySocialLogin")
	}
	if user == nil {
		user, err = lib.CreateUserFromSocialLogin(lib.AUTH_PROVIDER_MASTODON, userDetails.ID)
		if err != nil {
			return nil, errors.Wrap(err, "(BuildMastodonResponse) CreateUserFromSocialLogin")
		}
	}

	err = lib.SaveMastodonAccessToken(*user.Id, tokens.AccessToken, instanceDomain)

	tokenStr, err := MintJwt(*user.Id, lib.AUTH_PROVIDER_MASTODON)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildMastodonResponse) mint jwt")
	}

	rv := ResponseBody{
		AccessToken:     *tokenStr,
		Id:              fmt.Sprint(*user.Id),
		Name:            userDetails.DisplayName,
		ProfileImageUrl: userDetails.Avatar,
		Username:        userDetails.Username,
		Service:         "mastodon",
	}
	return &rv, nil
}

func MintJwt(userId int64, serviceId int) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    fmt.Sprint(userId),
		"service_id": fmt.Sprint(serviceId),
		"nbf":        time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, errors.Wrap(err, "(MintJwt) token.SignedString")
	}
	return &tokenString, err
}
