package alchemyapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextGetRankedNamedEntities(t *testing.T) {
	c := client()
	text := "The more things change... Yes, I'm inclined to agree, especially with regards to the historical relationship between stock prices and bond yields. "
	in := c.NewTextGetRankedNamedEntities(text)
	out, err := c.TextGetRankedNamedEntities(in)
	assert.Nil(t, err)
	assert.Equal(t, "OK", out.Status, out.StatusInfo)
}
