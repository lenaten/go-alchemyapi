package alchemyapi

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

// URLGetTextInput request output.
type URLGetTextInput struct {
	APIKey     string `url:"apikey"`
	OutputMode string `url:"outputMode"`
	URL        string `url:"url"`
}

// URLGetTextOutput request output.
type URLGetTextOutput struct {
	Output
	Text string `json:"text"`
}

func (c *Client) NewURLGetTextInput(url string) *URLGetTextInput {
	return &URLGetTextInput{
		APIKey:     c.APIKey,
		OutputMode: "json",
		URL:        url,
	}
}

// URLGetText returns a URLGetTextOutput.
func (c *Client) URLGetText(in *URLGetTextInput) (out *URLGetTextOutput, err error) {
	v, _ := query.Values(in)
	body, err := c.call("/url/URLGetText", v.Encode())
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
