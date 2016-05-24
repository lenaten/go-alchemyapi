package alchemyapi

import (
	"encoding/json"
	"strings"
)

// URLGetCombinedDataInput request output.
type URLGetCombinedDataInput struct {
	URL     string `url:"url"`
	Extract string `url:"extract"`
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
		URL:     url,
		Extract: strings.Join(extract, ","),
	}
}

// URLGetCombinedData returns a URLGetCombinedDataOutput.
func (c *Client) URLGetCombinedData(in *URLGetCombinedDataInput) (out *URLGetCombinedDataOutput, err error) {
	body, err := c.call("/url/URLGetCombinedData", map[string]string{
		"url":     in.URL,
		"extract": in.Extract,
	})
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
