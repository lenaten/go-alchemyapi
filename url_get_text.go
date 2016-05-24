package alchemyapi

import "encoding/json"

// URLGetTextInput request output.
type URLGetTextInput struct {
	URL string `url:"url"`
}

// URLGetTextOutput request output.
type URLGetTextOutput struct {
	Output
	Text string `json:"text"`
}

func (c *Client) NewURLGetTextInput(url string) *URLGetTextInput {
	return &URLGetTextInput{
		URL: url,
	}
}

// URLGetText returns a URLGetTextOutput.
func (c *Client) URLGetText(in *URLGetTextInput) (out *URLGetTextOutput, err error) {
	// v, _ := query.Values(in)
	body, err := c.call("/url/URLGetText", map[string]string{
		"url": in.URL,
	})
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
