package alchemyapi

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client implements a AlchemyAPI client.
type Client struct {
	*Config
}

type Entity struct {
	Type          string        `json:"type"`
	Relevance     string        `json:"relevance"`
	Count         string        `json:"count"`
	Text          string        `json:"text"`
	Disambiguated Disambiguated `json:"disambiguated"`
}

type Disambiguated struct {
	SubType  []string `json:"subType"`
	Name     string   `json:"name"`
	Freebase string   `json:"freebase"`
	Website  string   `json:"website"`
}

type Output struct {
	Status         string `json:"status"`
	StatusInfo     string `json:"statusInfo"`
	WarningMessage string `json:"warningMessage"`
	Usage          string `json:"usage"`
	URL            string `json:"url"`
}

// New client.
func New(config *Config) *Client {
	c := &Client{Config: config}
	return c
}

// call rpc style endpoint.
func (c *Client) call(path string, in map[string]string) (io.ReadCloser, error) {

	query := url.Values{}
	query.Set("apikey", c.APIKey)
	query.Set("outputMode", "json")

	u := "http://gateway-a.watsonplatform.net/calls" + path + "?" + query.Encode()

	form := url.Values{}
	for key, value := range in {
		form.Add(key, value)
	}

	req, err := http.NewRequest("POST", u, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.PostForm = form
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r, _, err := c.do(req)
	return r, err
}

// perform the request.
func (c *Client) do(req *http.Request) (io.ReadCloser, int64, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	if res.StatusCode < 400 {
		return res.Body, res.ContentLength, err
	}

	defer res.Body.Close()

	e := &Error{
		Status:     http.StatusText(res.StatusCode),
		StatusCode: res.StatusCode,
	}

	kind := res.Header.Get("Content-Type")

	if strings.Contains(kind, "text/plain") {
		if b, err := ioutil.ReadAll(res.Body); err == nil {
			e.Summary = string(b)
			return nil, 0, e
		}

		return nil, 0, err
	}

	if err := json.NewDecoder(res.Body).Decode(e); err != nil {
		return nil, 0, err
	}

	return nil, 0, e
}
