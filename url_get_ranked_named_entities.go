package alchemyapi

import "encoding/json"

// URLGetRankedNamedEntitiesInput request output.
type URLGetRankedNamedEntitiesInput struct {
	URL string `url:"url"`
}

// URLGetRankedNamedEntitiesOutput request output.
type URLGetRankedNamedEntitiesOutput struct {
	Output
	Language string   `json:"language"`
	Entities []Entity `json:"entities"`
}

func (c *Client) NewURLGetRankedNamedEntitiesInput(url string) *URLGetRankedNamedEntitiesInput {
	return &URLGetRankedNamedEntitiesInput{
		URL: url,
	}
}

// URLGetRankedNamedEntities returns a URLGetRankedNamedEntitiesOutput.
func (c *Client) URLGetRankedNamedEntities(in *URLGetRankedNamedEntitiesInput) (out *URLGetRankedNamedEntitiesOutput, err error) {
	body, err := c.call("/url/URLGetRankedNamedEntities", map[string]string{
		"url": in.URL,
	})
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
