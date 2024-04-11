package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Listing struct {
	Title string
	Price string
	Location string
	Link string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	listings := []Listing{}

	c.OnHTML("li.cl-static-search-result", func(e *colly.HTMLElement) {
	
		title := e.ChildText("div.title")
		price := e.ChildText("div.price")
		location := e.ChildText("div.location")
		link := e.ChildAttr("a", "href")
		
		listing := Listing{title, price, location, link}
		listings = append(listings, listing)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(listings)
	})


	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://vancouver.craigslist.org/search/apa")
}
