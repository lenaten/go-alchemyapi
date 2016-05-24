package alchemyapi

import "encoding/json"

type Authors struct {
	Names []string `json:"names"`
}

// URLGetAuthorsInput request output.
type URLGetAuthorsInput struct {
	URL string `url:"url"`
}

// URLGetAuthorsOutput request output.
type URLGetAuthorsOutput struct {
	Output
	Authors Authors `json:"authors"`
}

func (c *Client) NewURLGetAuthorsInput(url string) *URLGetAuthorsInput {
	return &URLGetAuthorsInput{
		URL: url,
	}
}

// URLGetAuthors returns a URLGetAuthorsOutput.
func (c *Client) URLGetAuthors(in *URLGetAuthorsInput) (out *URLGetAuthorsOutput, err error) {
	// v, _ := query.Values(in)
	body, err := c.call("/url/URLGetAuthors", map[string]string{
		"url": in.URL,
	})
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
