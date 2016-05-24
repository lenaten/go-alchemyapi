package alchemyapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLGetAuthors(t *testing.T) {
	c := client()
	url := "http://www.dailymail.co.uk/sciencetech/article-2355833/Apples-iPhone-5-hated-handset--majority-people-love-Samsung-Galaxy-S4-study-finds.html"
	in := c.NewURLGetAuthorsInput(url)
	out, err := c.URLGetAuthors(in)
	assert.Nil(t, err)
	assert.Equal(t, "OK", out.Status, out.StatusInfo)
	assert.Equal(t, url, out.URL)
	authors := []string{"Victoria Woollaston"}
	assert.Equal(t, out.Authors.Names, authors)
}
