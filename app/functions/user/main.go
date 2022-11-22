package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	utils "github.com/bmorrisondev/go-utils"
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
	Code string `json:"code"`
}

func Post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body RequestBody
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return utils.ErrorResponse(err, "json.Unmarshal")
	}

	data := url.Values{
		"code":          {body.Code},
		"grant_type":    {"authorization_code"},
		"client_id":     {"UFBHNHNySjVId2VLWWhSMTIyQ1o6MTpjaQ"},
		"redirect_uri":  {"http://localhost:8888/auth"},
		"code_verifier": {"challenge"},
	}

	req, err := http.NewRequest("POST", "https://api.twitter.com/2/oauth2/token", strings.NewReader(data.Encode()))
	req.SetBasicAuth(os.Getenv("TWITTER_CLIENT_ID"), os.Getenv("TWITTER_CLIENT_SECRET"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return utils.ErrorResponse(err, "client.Do")
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	jstr := string(bodyText)

	// resp, err := http.PostForm(, data)
	// if err != nil {
	// 	return utils.ErrorResponse(err, "http.PostForm")
	// }
	// defer resp.Body.Close()

	// var res map[string]interface{}
	// json.NewDecoder(resp.Body).Decode(&res)

	// jstr, err := utils.ConvertToJsonString(res)
	// if err != nil {
	// 	return utils.ErrorResponse(err, "utils.ConverToJsonString")
	// }
	return utils.OkResponse(&jstr)
}

func main() {
	lambda.Start(handler)
}
