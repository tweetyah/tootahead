package main

import (
	"encoding/json"
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

	log.Println("twitterAuthResp", twitterAuthResp)

	userDetails, err := lib.GetTwitterUserDetails(twitterAuthResp.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildTwitterResponse) GetTwitterUserDetails")
	}

	log.Println("userDetails", userDetails)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"twitter:access_token":      twitterAuthResp.AccessToken,
		"twitter:refresh_token":     twitterAuthResp.RefreshToken,
		"twitter:expires_in":        twitterAuthResp.ExpiresIn,
		"twitter:user_id":           userDetails.Data.Id,
		"twitter:username":          userDetails.Data.Username,
		"twitter:profile_image_url": userDetails.Data.ProfileImageUrl,
		"twitter:name":              userDetails.Data.Name,
		"nbf":                       time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, errors.Wrap(err, "(BuildTwitterResponse) token.SignedString")
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

	rv := ResponseBody{
		AccessToken:     tokenString,
		Id:              userDetails.Data.Id,
		Name:            userDetails.Data.Name,
		ProfileImageUrl: userDetails.Data.ProfileImageUrl,
		Username:        userDetails.Data.Username,
	}
	return &rv, nil
}

func BuildMastodonResponse(instanceDomain, code string) (*ResponseBody, error) {
	tokens, err := lib.GetMastodonTokens(instanceDomain, code)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildMastodonResponse) GetMastodonTokens")
	}

	userDetails, err := lib.GetMastodonUserDetails(instanceDomain, tokens.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "(BuildMastodonResponse) GetMastodonUserDetails")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mastodon:access_token":      tokens.AccessToken,
		"mastodon:user_id":           userDetails.ID,
		"mastodon:username":          userDetails.Username,
		"mastodon:profile_image_url": userDetails.Avatar,
		"mastodon:name":              userDetails.DisplayName,
		"nbf":                        time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, errors.Wrap(err, "(BuildMastodonResponse) token.SignedString")
	}

	rv := ResponseBody{
		AccessToken:     tokenString,
		Id:              userDetails.ID,
		Name:            userDetails.DisplayName,
		ProfileImageUrl: userDetails.Avatar,
		Username:        userDetails.Username,
	}
	return &rv, nil
}
