package alchemyapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextGetRankedNamedEntities(t *testing.T) {
	c := client()
	text := "The more things change... Yes, I'm inclined to agree, especially with regards to the historical relationship between stock prices and bond yields. The two have generally traded together, rising during periods of economic growth and falling during periods of contraction. Consider the period from 1998 through 2010, during which the U.S. economy experienced two expansions as well as two recessions: Then central banks came to the rescue. Fed Chairman Ben Bernanke led from Washington with the help of the bank's current $3.6T balance sheet. He's accompanied by Mario Draghi at the European Central Bank and an equally forthright Shinzo Abe in Japan. Their coordinated monetary expansion has provided all the sugar needed for an equities moonshot, while they vowed to hold global borrowing costs at record lows."
	in := c.NewTextGetRankedNamedEntities(text)
	out, err := c.TextGetRankedNamedEntities(in)
	assert.Nil(t, err)
	assert.Equal(t, "OK", out.Status)
}
