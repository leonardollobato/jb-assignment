package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		// colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
		colly.AllowedDomains("news.ycombinator.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("span.titleline > a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))

		c.ta
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://news.ycombinator.com
	c.Visit("https://news.ycombinator.com/")
}