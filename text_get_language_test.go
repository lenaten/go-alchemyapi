package alchemyapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextGetLanguage(t *testing.T) {
	c := client()
	text := "The more things change"
	in := c.NewTextGetLanguage(text)
	out, err := c.TextGetLanguage(in)
	assert.Nil(t, err)
	assert.Equal(t, "OK", out.Status, out.StatusInfo)
	assert.Equal(t, "english", out.Language)
	assert.Equal(t, "en", out.ISO6391)
}
