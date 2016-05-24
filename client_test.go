package alchemyapi

import "github.com/segmentio/go-env"

func client() *Client {
	apiKey := env.MustGet("ALCHEMYAPI_API_KEY")
	return New(NewConfig(apiKey))
}
