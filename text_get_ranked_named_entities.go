package alchemyapi

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

// TextGetRankedNamedEntitiesInput request output.
type TextGetRankedNamedEntitiesInput struct {
	APIKey     string `url:"apikey"`
	OutputMode string `url:"outputMode"`
	Text       string `url:"text"`
}

// TextGetRankedNamedEntitiesOutput request output.
type TextGetRankedNamedEntitiesOutput struct {
	Status     string   `json:"status"`
	StatusInfo string   `json:"statusInfo"`
	Usage      string   `json:"usage"`
	Language   string   `json:"language"`
	Entities   []Entity `json:"entities"`
}

func (c *Client) NewTextGetRankedNamedEntities(text string) *TextGetRankedNamedEntitiesInput {
	return &TextGetRankedNamedEntitiesInput{
		APIKey:     c.APIKey,
		OutputMode: "json",
		Text:       text,
	}
}

// TextGetRankedNamedEntities returns a TextGetRankedNamedEntitiesOutput.
func (c *Client) TextGetRankedNamedEntities(in *TextGetRankedNamedEntitiesInput) (out *TextGetRankedNamedEntitiesOutput, err error) {
	v, _ := query.Values(in)
	body, err := c.call("/text/TextGetRankedNamedEntities", v.Encode())
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
