package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func GetTwitterTokens(code string) (*TwitterAuthResponse, error) {
	data := url.Values{
		"code":          {code},
		"grant_type":    {"authorization_code"},
		"client_id":     {os.Getenv("VITE_TWITTER_CLIENT_ID")},
		"redirect_uri":  {os.Getenv("VITE_REDIRECT_URI")},
		"code_verifier": {"challenge"},
	}

	req, err := http.NewRequest("POST", "https://api.twitter.com/2/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "(GetTwitterTokens) http.NewRequest")
	}
	req.SetBasicAuth(os.Getenv("VITE_TWITTER_CLIENT_ID"), os.Getenv("TWITTER_CLIENT_SECRET"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "(GetTwitterTokens) client.Do")
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "(GetTwitterTokens) ioutil.ReadAll")
	}

	var twitterAuthResp TwitterAuthResponse
	err = json.Unmarshal([]byte(bodyText), &twitterAuthResp)
	if err != nil {
		return nil, errors.Wrap(err, "(GetTwitterTokens) json.Unmarshal")
	}
	return &twitterAuthResp, nil
}

type TwitterAuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

func GetTwitterUserDetails(token string) (*TwitterUserResponse, error) {
	req, err := http.NewRequest("GET", "https://api.twitter.com/2/users/me?user.fields=profile_image_url", nil)
	if err != nil {
		return nil, errors.Wrap(err, "(GetTwitterUserDetails) http.NewRequest")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "(GetTwitterUserDetails) client.Do")
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "(GetTwitterUserDetail) ioutil.ReadAll")
	}

	var response TwitterUserResponse
	err = json.Unmarshal([]byte(bodyText), &response)
	if err != nil {
		return nil, errors.Wrap(err, "(GetTwitterUserDetails) json.Unmashal")
	}
	return &response, nil
}

type TwitterUserResponse struct {
	Data struct {
		Id              string `json:"id"`
		Name            string `json:"name"`
		ProfileImageUrl string `json:"profile_image_url"`
		Username        string `json:"username"`
	} `json:"data"`
}

type SendTweetRequest struct {
	Text *string `json:"text"`
}

type SendTweetResponse struct {
	Data struct {
		Id *string `json:"id"`
	} `json:"data"`
	Title  *string `json:"title"`
	Type   *string `json:"type"`
	Status *int    `json:"status"`
	Detail *string `json:"detail"`
}

type SendTweetResults struct {
	SentId    *int64
	IsSuccess bool
	Status    *string
}

// Sends the tweet and returns the id of the tweet sent
func SendTweet(text string, token string) (*SendTweetResults, error) {
	requestBody := SendTweetRequest{
		Text: &text,
	}
	jbytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, errors.Wrap(err, "(SendTweet) json.Marshal")
	}

	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer(jbytes))
	if err != nil {
		return nil, errors.Wrap(err, "(SendTweet) http.NewRequest")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "(SendTweet) client.Do")
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "(SendTweet) ioutil.ReadAll")
	}

	var response SendTweetResponse
	err = json.Unmarshal([]byte(bodyText), &response)
	if err != nil {
		return nil, errors.Wrap(err, "(SendTweet) json.Unmarshall")
	}

	if response.Status != nil {
		rawResponse := string(bodyText)
		results := SendTweetResults{
			Status: &rawResponse,
		}
		return &results, nil
	} else {
		idNum, err := strconv.ParseInt(*response.Data.Id, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, "(SendTweet) strconv.ParseInt")
		}
		results := SendTweetResults{
			SentId:    &idNum,
			IsSuccess: true,
		}
		return &results, nil
	}
}

func GetTwitterTokensViaRefresh(refreshToken string) (*TwitterAuthResponse, error) {
	data := url.Values{
		"refresh_token": {refreshToken},
		"grant_type":    {"refresh_token"},
		"client_id":     {os.Getenv("VITE_TWITTER_CLIENT_ID")},
	}

	req, err := http.NewRequest("POST", "https://api.twitter.com/2/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "(GetTwitterTokensViaRefresh) http.NewRequest")
	}
	req.SetBasicAuth(os.Getenv("VITE_TWITTER_CLIENT_ID"), os.Getenv("TWITTER_CLIENT_SECRET"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "(GetTwitterTokensViaRefresh) client.Do")
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "(GetTwitterTokensViaRefresh) ioutil.ReadAll")
	}

	var twitterAuthResp TwitterAuthResponse
	err = json.Unmarshal([]byte(bodyText), &twitterAuthResp)
	if err != nil {
		return nil, errors.Wrap(err, "(GetTwitterTokensViaRefresh) json.Unmarshal")
	}
	return &twitterAuthResp, nil

}
