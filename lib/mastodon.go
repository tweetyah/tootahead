package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func GetMastodonTokens(instanceDomain, code, clientId, clientSecret string) (*MastodonAuthResponse, error) {
	data := url.Values{
		"code":          {code},
		"grant_type":    {"authorization_code"},
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"redirect_uri":  {os.Getenv("VITE_REDIRECT_URI")},
		"scope":         {"read write follow"},
	}

	mastodonUrl := fmt.Sprintf("https://%v/oauth/token", instanceDomain)

	req, err := http.NewRequest("POST", mastodonUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "(GetMastodonTokens) http.NewRequest")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "(GetMastodonTokens) client.Do")
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "(GetMastodonTokens) ioutil.ReadAll")
	}

	var authResp MastodonAuthResponse
	err = json.Unmarshal([]byte(bodyText), &authResp)
	if err != nil {
		return nil, errors.Wrap(err, "(GetMastodonTokens) json.Unmarshal")
	}
	return &authResp, nil
}

type MastodonAuthResponse struct {
	AccessToken string `json:"access_token"`
	CreatedAt   int    `json:"created_at"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func GetMastodonUserDetails(instanceDomain, token string) (*MastodonGetUserResponse, error) {
	mastodonUrl := fmt.Sprintf("https://%v/api/v1/accounts/verify_credentials", instanceDomain)
	req, err := http.NewRequest("GET", mastodonUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, "(GetMastodonUserDetails) http.NewRequest")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "(GetMastodonUserDetails) client.Do")
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "(GetMastodonUserDetails) ioutil.ReadAll")
	}

	var response MastodonGetUserResponse
	err = json.Unmarshal([]byte(bodyText), &response)
	if err != nil {
		return nil, errors.Wrap(err, "(GetMastodonUserDetails) json.Unmashal")
	}
	return &response, nil
}

type MastodonGetUserResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Acct        string `json:"acct"`
	DisplayName string `json:"display_name"`
	Locked      bool   `json:"locked"`
	Bot         bool   `json:"bot"`
	// CreatedAt      time.Time `json:"created_at"`
	Note           string `json:"note"`
	URL            string `json:"url"`
	Avatar         string `json:"avatar"`
	AvatarStatic   string `json:"avatar_static"`
	Header         string `json:"header"`
	HeaderStatic   string `json:"header_static"`
	FollowersCount int    `json:"followers_count"`
	FollowingCount int    `json:"following_count"`
	StatusesCount  int    `json:"statuses_count"`
	// LastStatusAt   time.Time `json:"last_status_at"`
	Source struct {
		Privacy   string `json:"privacy"`
		Sensitive bool   `json:"sensitive"`
		Language  string `json:"language"`
		Note      string `json:"note"`
		Fields    []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
			// VerifiedAt time.Time `json:"verified_at"`
		} `json:"fields"`
		FollowRequestsCount int `json:"follow_requests_count"`
	} `json:"source"`
	Emojis []struct {
		Shortcode       string `json:"shortcode"`
		URL             string `json:"url"`
		StaticURL       string `json:"static_url"`
		VisibleInPicker bool   `json:"visible_in_picker"`
	} `json:"emojis"`
	Fields []struct {
		Name       string    `json:"name"`
		Value      string    `json:"value"`
		VerifiedAt time.Time `json:"verified_at"`
	} `json:"fields"`
}

func SendMastodonPost(instanceDomain string, text string, accessToken string) (*SendMastodonPostResults, error) {
	data := url.Values{
		"status": {text},
	}
	mastodonUrl := fmt.Sprintf("https://%v/api/v1/statuses", instanceDomain)
	opts := FetchOptions{
		Method:   "POST",
		FormData: &data,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", accessToken),
		},
	}
	res, err := Fetch(mastodonUrl, &opts)
	if err != nil {
		return nil, errors.Wrap(err, "(SendMastodonPost) fetch")
	}

	var resp SendMastodonPostResults
	err = res.MarshalJson(&resp)
	if err != nil {
		return nil, errors.Wrap(err, "(SendMastodonPost) marshal json")
	}
	return &resp, nil
}

type SendMastodonPostResults struct {
	ID                 string        `json:"id"`
	CreatedAt          time.Time     `json:"created_at"`
	InReplyToID        interface{}   `json:"in_reply_to_id"`
	InReplyToAccountID interface{}   `json:"in_reply_to_account_id"`
	Sensitive          bool          `json:"sensitive"`
	SpoilerText        string        `json:"spoiler_text"`
	Visibility         string        `json:"visibility"`
	Language           string        `json:"language"`
	URI                string        `json:"uri"`
	URL                string        `json:"url"`
	RepliesCount       int           `json:"replies_count"`
	ReblogsCount       int           `json:"reblogs_count"`
	FavouritesCount    int           `json:"favourites_count"`
	EditedAt           interface{}   `json:"edited_at"`
	Favourited         bool          `json:"favourited"`
	Reblogged          bool          `json:"reblogged"`
	Muted              bool          `json:"muted"`
	Bookmarked         bool          `json:"bookmarked"`
	Pinned             bool          `json:"pinned"`
	Content            string        `json:"content"`
	Filtered           []interface{} `json:"filtered"`
	Reblog             interface{}   `json:"reblog"`
	Application        struct {
		Name    string `json:"name"`
		Website string `json:"website"`
	} `json:"application"`
	Account struct {
		ID             string    `json:"id"`
		Username       string    `json:"username"`
		Acct           string    `json:"acct"`
		DisplayName    string    `json:"display_name"`
		Locked         bool      `json:"locked"`
		Bot            bool      `json:"bot"`
		Discoverable   bool      `json:"discoverable"`
		Group          bool      `json:"group"`
		CreatedAt      time.Time `json:"created_at"`
		Note           string    `json:"note"`
		URL            string    `json:"url"`
		Avatar         string    `json:"avatar"`
		AvatarStatic   string    `json:"avatar_static"`
		Header         string    `json:"header"`
		HeaderStatic   string    `json:"header_static"`
		FollowersCount int       `json:"followers_count"`
		FollowingCount int       `json:"following_count"`
		StatusesCount  int       `json:"statuses_count"`
		LastStatusAt   string    `json:"last_status_at"`
		Noindex        bool      `json:"noindex"`
		Emojis         []struct {
			Shortcode       string `json:"shortcode"`
			URL             string `json:"url"`
			StaticURL       string `json:"static_url"`
			VisibleInPicker bool   `json:"visible_in_picker"`
		} `json:"emojis"`
		Fields []struct {
			Name       string      `json:"name"`
			Value      string      `json:"value"`
			VerifiedAt interface{} `json:"verified_at"`
		} `json:"fields"`
	} `json:"account"`
	MediaAttachments []interface{} `json:"media_attachments"`
	Mentions         []interface{} `json:"mentions"`
	Tags             []interface{} `json:"tags"`
	Emojis           []interface{} `json:"emojis"`
	Card             interface{}   `json:"card"`
	Poll             interface{}   `json:"poll"`
	Error            *string       `json:"error"`
}

func RegisterMastodonApp(domain string) (*MastodonAppRegistration, error) {
	formData := url.Values{
		"client_name":   {"Tweetyah"},
		"redirect_uris": {os.Getenv("VITE_REDIRECT_URI")},
		"scopes":        {"read write follow"},
		"website":       {"https://tweetyah.com"},
	}
	url := fmt.Sprintf("https://%v/api/v1/apps", domain)
	opts := FetchOptions{
		Method:   "POST",
		FormData: &formData,
	}
	res, err := Fetch(url, &opts)
	if err != nil {
		return nil, errors.Wrap(err, "(RegisterMastodonApp) fetch")
	}

	var appRegistration MastodonAppRegistration
	err = res.MarshalJson(&appRegistration)
	if err != nil {
		return nil, errors.Wrap(err, "(RegisterMastodonApp) marshal json")
	}

	return &appRegistration, nil
}

type MastodonAppRegistration struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Website      string `json:"website"`
	RedirectURI  string `json:"redirect_uri"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	VapidKey     string `json:"vapid_key"`
}
