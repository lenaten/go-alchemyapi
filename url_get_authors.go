package alchemyapi

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

type Authors struct {
	Names []string `json:"names"`
}

// URLGetAuthorsInput request output.
type URLGetAuthorsInput struct {
	APIKey     string `url:"apikey"`
	OutputMode string `url:"outputMode"`
	URL        string `url:"url"`
}

// URLGetAuthorsOutput request output.
type URLGetAuthorsOutput struct {
	Status     string  `json:"status"`
	StatusInfo string  `json:"statusInfo"`
	Usage      string  `json:"usage"`
	URL        string  `json:"url"`
	Authors    Authors `json:"authors"`
}

func (c *Client) NewURLGetAuthorsInput(url string) *URLGetAuthorsInput {
	return &URLGetAuthorsInput{
		APIKey:     c.APIKey,
		OutputMode: "json",
		URL:        url,
	}
}

// URLGetAuthors returns a URLGetAuthorsOutput.
func (c *Client) URLGetAuthors(in *URLGetAuthorsInput) (out *URLGetAuthorsOutput, err error) {
	v, _ := query.Values(in)
	body, err := c.call("/url/URLGetAuthors", v.Encode())
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
