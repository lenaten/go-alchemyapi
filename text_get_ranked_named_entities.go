package alchemyapi

import "encoding/json"

// TextGetRankedNamedEntitiesInput request output.
type TextGetRankedNamedEntitiesInput struct {
	Text string `url:"text"`
}

// TextGetRankedNamedEntitiesOutput request output.
type TextGetRankedNamedEntitiesOutput struct {
	Output
	Language string   `json:"language"`
	Entities []Entity `json:"entities"`
}

func (c *Client) NewTextGetRankedNamedEntities(text string) *TextGetRankedNamedEntitiesInput {
	return &TextGetRankedNamedEntitiesInput{
		Text: text,
	}
}

// TextGetRankedNamedEntities returns a TextGetRankedNamedEntitiesOutput.
func (c *Client) TextGetRankedNamedEntities(in *TextGetRankedNamedEntitiesInput) (out *TextGetRankedNamedEntitiesOutput, err error) {
	body, err := c.call("/text/TextGetRankedNamedEntities", map[string]string{
		"text": in.Text,
	})
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
