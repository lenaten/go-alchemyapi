package alchemyapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLGetCombinedData(t *testing.T) {
	c := client()
	url := "http://www.dailymail.co.uk/sciencetech/article-2355833/Apples-iPhone-5-hated-handset--majority-people-love-Samsung-Galaxy-S4-study-finds.html"
	extract := []string{"author", "entity"}
	in := c.NewURLGetCombinedDataInput(url, extract)
	out, err := c.URLGetCombinedData(in)
	assert.Nil(t, err)
	assert.Equal(t, "OK", out.Status)
	assert.Equal(t, url, out.URL)
	assert.Equal(t, "Victoria Woollaston", out.Author)
}
