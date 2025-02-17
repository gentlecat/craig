package craig

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInvalidListing(t *testing.T) {
	listing, err := ParseListing(strings.NewReader(""))
	assert.Error(t, errInvalidListing, err)
	assert.Nil(t, listing)
}

func TestParseEmptyListing(t *testing.T) {
	file, err := os.Open("./fixtures/empty_listing.html")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	listing, err := ParseListing(file)
	assert.NoError(t, err)
	assert.NotNil(t, listing)
	assert.Equal(t, "6569794207", listing.ID)
}

func TestParseListing(t *testing.T) {
	file, err := os.Open("./fixtures/listing.html")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	listing, err := ParseListing(file)
	assert.NoError(t, err)
	assert.NotNil(t, listing)
	assert.Equal(t, "7128708803", listing.ID)
	assert.Equal(t, "1950 chrysler", listing.Title)
	assert.Equal(t, "https://chicago.craigslist.org/nwc/cto/d/wood-dale-1950-chrysler/7128708803.html", listing.URL)
	assert.NotEmpty(t, listing.Description)
	// assert.Equal(t, uint(3000), listing.Price)
	assert.Equal(t, 14, len(listing.Images))
	assert.NotNil(t, listing.PostedAt)
	assert.NotNil(t, listing.UpdatedAt)
	assert.Equal(t, 41.9602, listing.Location.Lat)
	assert.Equal(t, -87.981, listing.Location.Lng)

	attrs := map[string]string{
		"fuel": "gas",
	}
	for k, v := range attrs {
		assert.Equal(t, v, listing.Attributes[k])
	}
}
