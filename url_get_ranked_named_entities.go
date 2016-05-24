package alchemyapi

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

// URLGetRankedNamedEntitiesInput request output.
type URLGetRankedNamedEntitiesInput struct {
	APIKey     string `url:"apikey"`
	OutputMode string `url:"outputMode"`
	URL        string `url:"url"`
}

// URLGetRankedNamedEntitiesOutput request output.
type URLGetRankedNamedEntitiesOutput struct {
	Status     string   `json:"status"`
	StatusInfo string   `json:"statusInfo"`
	Usage      string   `json:"usage"`
	URL        string   `json:"url"`
	Language   string   `json:"language"`
	Entities   []Entity `json:"entities"`
}

func (c *Client) NewURLGetRankedNamedEntitiesInput(url string) *URLGetRankedNamedEntitiesInput {
	return &URLGetRankedNamedEntitiesInput{
		APIKey:     c.APIKey,
		OutputMode: "json",
		URL:        url,
	}
}

// URLGetRankedNamedEntities returns a URLGetRankedNamedEntitiesOutput.
func (c *Client) URLGetRankedNamedEntities(in *URLGetRankedNamedEntitiesInput) (out *URLGetRankedNamedEntitiesOutput, err error) {
	v, _ := query.Values(in)
	body, err := c.call("/url/URLGetRankedNamedEntities", v.Encode())
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
