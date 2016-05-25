package alchemyapi

import "encoding/json"

// TextGetLanguageInput request output.
type TextGetLanguageInput struct {
	Text string `url:"text"`
}

// TextGetLanguageOutput request output.
type TextGetLanguageOutput struct {
	Output
	Language string `json:"language"`
	ISO6391  string `json:"iso-639-1"`
}

func (c *Client) NewTextGetLanguage(text string) *TextGetLanguageInput {
	return &TextGetLanguageInput{
		Text: text,
	}
}

// TextGetLanguage returns a TextGetLanguageOutput.
func (c *Client) TextGetLanguage(in *TextGetLanguageInput) (out *TextGetLanguageOutput, err error) {
	body, err := c.call("/text/TextGetLanguage", map[string]string{
		"text": in.Text,
	})
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
