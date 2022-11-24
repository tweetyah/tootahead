package lib

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type FetchOptions struct {
	UrlBase  *string
	Method   string
	Headers  map[string]string
	Body     *string
	FormData *url.Values
}

type FetchResponse struct {
	RawResponse *http.Response
}

func (r *FetchResponse) MarshalJson(out interface{}) error {
	b, err := io.ReadAll(r.RawResponse.Body)
	if err != nil {
		return errors.Wrap(err, "(Json) failed to read body")
	}
	err = json.Unmarshal(b, &out)
	if err != nil {
		return errors.Wrap(err, "(Json) unmarshal json")
	}
	return nil
}

func (r *FetchResponse) Text() (*string, error) {
	b, err := io.ReadAll(r.RawResponse.Body)
	if err != nil {
		return nil, errors.Wrap(err, "(Text) failed to read body")
	}
	str := string(b)
	return &str, nil
}

func Fetch(url string, options *FetchOptions) (*FetchResponse, error) {
	if options == nil {
		options = &FetchOptions{
			Method: "GET",
		}
	}

	var req *http.Request
	if options.UrlBase != nil {
		url = *options.UrlBase + url
	}

	if options.Body != nil {
		bodyReader := strings.NewReader(*options.Body)
		r, err := http.NewRequest(options.Method, url, bodyReader)
		if err != nil {
			return nil, errors.Wrap(err, "(Fetch) create request")
		}
		req = r
	} else if options.FormData != nil {
		data := *options.FormData
		r, err := http.NewRequest(options.Method, url, strings.NewReader(data.Encode()))
		if err != nil {
			return nil, errors.Wrap(err, "(Fetch) create request")
		}
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req = r
	} else {
		r, err := http.NewRequest(options.Method, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "(Fetch) create request")
		}
		req = r
	}

	if options.Headers != nil {
		for key, val := range options.Headers {
			req.Header.Add(key, val)
		}
	}

	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "(Fetch) execute request")
	}
	fr := FetchResponse{
		RawResponse: res,
	}

	return &fr, nil
}
