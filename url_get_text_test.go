package alchemyapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLGetText(t *testing.T) {
	c := client()
	url := "http://www.dailymail.co.uk/sciencetech/article-2355833/Apples-iPhone-5-hated-handset--majority-people-love-Samsung-Galaxy-S4-study-finds.html"
	in := c.NewURLGetTextInput(url)
	out, err := c.URLGetText(in)
	assert.Nil(t, err)
	assert.Equal(t, "OK", out.Status)
	assert.Equal(t, url, out.URL)
}
