package scraper

import (
	"fmt"
	"net/url"

	"github.com/gocolly/colly/v2"
)

type Scraper struct {
}

func (s *Scraper) ScrapeUrl(siteUrl, htmlSelector string, parser func(e *colly.HTMLElement) (interface{}, error)) ([]interface{}, error) {
	var result []interface{}
	var err error

	fmt.Println("Scraping website", siteUrl)

	u, err := url.Parse(siteUrl)
	if err != nil {
		return nil, err
	}

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(u.Host),
	)

	// c.OnHTML("*", func(e *colly.HTMLElement) {
	// 	fmt.Println(e)
	// })

	c.OnHTML(htmlSelector, func(e *colly.HTMLElement) {
		if parsed, err := parser(e); err == nil {
			result = append(result, parsed)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	err = c.Visit(u.String())
	if err != nil {
		return nil, err
	}

	return result, nil
}
