package craig

import (
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	homeURL = "https://www.craigslist.org/about/sites"
)

type Site struct {
	ID         string
	Name       string
	RegionID   string
	RegionName string
	URL        string
}

// Sites returns all US-based sites
func Sites() ([]Site, error) {
	resp, err := http.Get(homeURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ParseSites(resp.Body)
}

// ParseSites extracts all sites from the reader
func ParseSites(reader io.Reader) ([]Site, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	sites := []Site{}
	boxes := doc.Find("div.colmask").First().Find("div.box")
	lastRegionID := ""
	lastReionName := ""

	boxes.Each(func(i int, el *goquery.Selection) {
		el.Children().Each(func(idx int, boxEl *goquery.Selection) {
			if idx%2 == 0 {
				lastRegionID = strings.ToLower(boxEl.Text())
				lastReionName = strings.Title(strings.ToLower(boxEl.Text()))
			} else {
				boxEl.Find("li > a").Each(func(linkIdx int, linkEl *goquery.Selection) {
					site := parseSiteLink(linkEl)
					if site != nil {
						site.RegionID = lastRegionID
						site.RegionName = lastReionName
						sites = append(sites, *site)
					}
				})
			}
		})
	})

	return sites, nil
}

func parseSiteLink(link *goquery.Selection) *Site {
	href := link.AttrOr("href", "")
	matches := reSiteID.FindStringSubmatch(href)

	if len(matches) == 0 {
		return nil
	}

	return &Site{
		ID:   matches[1],
		Name: strings.Title(strings.ToLower(link.Text())),
		URL:  strings.TrimRight(href, "/"),
	}
}
