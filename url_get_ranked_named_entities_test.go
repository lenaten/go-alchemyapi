package alchemyapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLGetRankedNamedEntities(t *testing.T) {
	c := client()
	url := "http://www.dailymail.co.uk/sciencetech/article-2355833/Apples-iPhone-5-hated-handset--majority-people-love-Samsung-Galaxy-S4-study-finds.html"
	in := c.NewURLGetRankedNamedEntitiesInput(url)
	out, err := c.URLGetRankedNamedEntities(in)
	assert.Equal(t, err, nil)
	assert.Equal(t, out.Status, "OK")
	assert.Equal(t, out.URL, url)
}
