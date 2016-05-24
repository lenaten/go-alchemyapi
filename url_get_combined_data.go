package alchemyapi

import (
	"encoding/json"
	"strings"

	"github.com/google/go-querystring/query"
)

// URLGetCombinedDataInput request output.
type URLGetCombinedDataInput struct {
	APIKey     string `url:"apikey"`
	OutputMode string `url:"outputMode"`
	URL        string `url:"url"`
	Extract    string `url:"extract"`
}

// URLGetCombinedDataOutput request output.
type URLGetCombinedDataOutput struct {
	Status     string   `json:"status"`
	StatusInfo string   `json:"statusInfo"`
	Usage      string   `json:"usage"`
	URL        string   `json:"url"`
	Language   string   `json:"language"`
	Author     string   `json:"author"`
	Entities   []Entity `json:"entities"`
}

func (c *Client) NewURLGetCombinedDataInput(url string, extract []string) *URLGetCombinedDataInput {
	return &URLGetCombinedDataInput{
		APIKey:     c.APIKey,
		OutputMode: "json",
		URL:        url,
		Extract:    strings.Join(extract, ","),
	}
}

// URLGetCombinedData returns a URLGetCombinedDataOutput.
func (c *Client) URLGetCombinedData(in *URLGetCombinedDataInput) (out *URLGetCombinedDataOutput, err error) {
	v, _ := query.Values(in)
	body, err := c.call("/url/URLGetCombinedData", v.Encode())
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
