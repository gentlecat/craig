package craig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSites(t *testing.T) {
	file, err := os.Open("./fixtures/sites.html")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	sites, err := ParseSites(file)
	assert.NoError(t, err)
	assert.Len(t, sites, 417)
}
